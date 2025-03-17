// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package query

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/esqlformat"
)

type metadata struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type esqlResponse struct {
	Columns []metadata `json:"columns"`
	Values  [][]any    `json:"values"`
}

// Helper takes a generic type T, a context.Context and an esql.Query request.
// Returns an array of T using the json.Unmarshaler of the type.
func Helper[T any](ctx context.Context, esqlQuery *Query) ([]T, error) {
	response, err := esqlQuery.
		Columnar(false).
		Format(esqlformat.Json).
		Header("x-elastic-client-meta", "h=qo").
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var eR esqlResponse
	err = json.Unmarshal(response, &eR)
	if err != nil {
		return nil, fmt.Errorf("cannot read ES|QL response: %w", err)
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteByte('[')
	for rowNum, row := range eR.Values {
		buf.WriteByte('{')
		for i := 0; i < len(row); i++ {
			buf.WriteString(`"` + eR.Columns[i].Name + `":`)
			data, err := json.Marshal(row[i])
			if err != nil {
				return nil, fmt.Errorf("error while parsing ES|QL response: %w", err)
			}
			buf.Write(data)
			if i != len(row)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')
		if rowNum != len(eR.Values)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')

	target := []T{}
	err = json.Unmarshal(buf.Bytes(), &target)
	if err != nil {
		return nil, fmt.Errorf("cannot deserialize ES|QL response: %w", err)
	}

	return target, nil
}

type EsqlIterator[T any] interface {
	Next() (*T, error)
	More() bool
}

type iterator[T any] struct {
	reader    []byte
	decoder   *json.Decoder
	keys      []string
	skipComma bool
}

func (d iterator[T]) More() bool {
	return d.decoder.More()
}

func (d iterator[T]) Next() (*T, error) {
	var t T
	var tmp []any

	if d.skipComma {
		d.decoder.Token()
	}

	err := d.decoder.Decode(&tmp)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}

	buf.WriteByte('{')
	for index, key := range d.keys {
		buf.WriteString(`"` + key + `":`)
		value, _ := json.Marshal(tmp[index])
		buf.Write(value)

		if index != len(d.keys)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')

	err = json.Unmarshal(buf.Bytes(), &t)
	if err != nil {
		return nil, err
	}

	d.skipComma = true
	return &t, nil
}

// Helper takes a generic type T, a context.Context and an esql.Query request
// buffer the response and provides an API to consume one item at a time.
func NewIteratorHelper[T any](ctx context.Context, query *Query) (EsqlIterator[T], error) {
	response, err := query.
		Columnar(false).
		Format(esqlformat.Json).
		Header("x-elastic-client-meta", "h=qo").
		Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	d := &iterator[T]{}
	d.reader, err = io.ReadAll(response.Body)
	d.decoder = json.NewDecoder(bytes.NewReader(d.reader))

	var metas []metadata
OUTER:
	for {
		t, err := d.decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		switch t {
		case "columns":
			err := d.decoder.Decode(&metas)
			if err != nil {
				return nil, err
			}
			for _, m := range metas {
				d.keys = append(d.keys, m.Name)
			}

		case "values":
			t, _ := d.decoder.Token()
			if t != json.Delim(91) {
				return nil, fmt.Errorf("cannot read response from ES|QL, expected ARRAY_START: %w", err)
			}
			break OUTER
		}
	}

	if err != nil {
		return nil, fmt.Errorf("cannot read response from ES|QL: %w", err)
	}

	return d, nil
}
