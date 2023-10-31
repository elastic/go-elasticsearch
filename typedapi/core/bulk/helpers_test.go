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

package bulk

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestBulk_CreateOp(t *testing.T) {
	indexName := "foo"

	type args struct {
		op  types.CreateOperation
		doc interface{}
	}
	tests := []struct {
		name    string
		args    args
		expects []byte
		wantErr bool
	}{
		{
			name: "struct doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{
					ID:   1,
					Name: "Foo",
				},
			},
			expects: []byte(`{"create":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "json doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				json.RawMessage(`{"id":1,"name":"Foo"}`),
			},
			expects: []byte(`{"create":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "truncated json doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				json.RawMessage(`{"id":1,"name}`),
			},
			expects: nil,
			wantErr: true,
		},
		{
			name: "bytes slice doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				[]byte(`{"id":1,"name":"Foo"}`),
			},
			expects: []byte(`{"create":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "bad bytes slice doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				[]byte(`{"id":1,"name","Foo}`),
			},
			expects: nil,
			wantErr: true,
		},
		{
			name: "not seralizable doc",
			args: args{
				types.CreateOperation{Index_: &indexName},
				math.NaN(),
			},
			expects: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.CreateOp(tt.args.op, tt.args.doc); (err != nil) != tt.wantErr {
				if r.buf.Len() == 0 {
					t.Errorf("buffer has not been reset")
				}
				t.Errorf("CreateOp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !bytes.Equal(r.buf.Bytes(), tt.expects) {
				t.Errorf("serialized buffer invalid\nexpected: %s\ngot: %s", tt.expects, r.buf)
			}
		})
	}
}

func TestBulk_DeleteOp(t *testing.T) {
	indexName := "foo"

	type args struct {
		op types.DeleteOperation
	}
	tests := []struct {
		name    string
		args    args
		expects []byte
		wantErr bool
	}{
		{
			name: "simple delete",
			args: args{
				op: types.DeleteOperation{
					Index_: &indexName,
				},
			},
			expects: []byte(`{"delete":{"_index":"foo"}}
`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.DeleteOp(tt.args.op); (err != nil) != tt.wantErr {
				if r.buf.Len() == 0 {
					t.Errorf("buffer has not been reset")
				}
				t.Errorf("DeleteOp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(r.buf.Bytes(), tt.expects) {
				t.Errorf("serialized buffer invalid\nexpected: %s\ngot: %s", tt.expects, r.buf)
			}
		})
	}
}

func TestBulk_IndexOp(t *testing.T) {
	indexName := "foo"

	type args struct {
		op  types.IndexOperation
		doc interface{}
	}
	tests := []struct {
		name    string
		args    args
		expects []byte
		wantErr bool
	}{
		{
			name: "struct doc",
			args: args{
				types.IndexOperation{Index_: &indexName},
				struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{
					ID:   1,
					Name: "Foo",
				},
			},
			expects: []byte(`{"index":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "json doc",
			args: args{
				types.IndexOperation{Index_: &indexName},
				json.RawMessage(`{"id":1,"name":"Foo"}`),
			},
			expects: []byte(`{"index":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "truncated json doc",
			args: args{
				types.IndexOperation{Index_: &indexName},
				json.RawMessage(`{"id":1,"name":"`),
			},
			expects: nil,
			wantErr: true,
		},
		{
			name: "bytes slice doc",
			args: args{
				types.IndexOperation{Index_: &indexName},
				[]byte(`{"id":1,"name":"Foo"}`),
			},
			expects: []byte(`{"index":{"_index":"foo"}}
{"id":1,"name":"Foo"}
`),
			wantErr: false,
		},
		{
			name: "bad bytes slice doc",
			args: args{
				types.IndexOperation{Index_: &indexName},
				[]byte(`{"id":1,"name","Foo}`),
			},
			expects: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.IndexOp(tt.args.op, tt.args.doc); (err != nil) != tt.wantErr {
				if r.buf.Len() == 0 {
					t.Errorf("buffer has not been reset")
				}
				t.Errorf("IndexOp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !bytes.Equal(r.buf.Bytes(), tt.expects) {
				t.Errorf("serialized buffer invalid\nexpected: %s\ngot: %s", tt.expects, r.buf)
			}
		})
	}
}

func TestBulk_UpdateOp(t *testing.T) {
	indexName := "foo"

	type args struct {
		op     types.UpdateOperation
		doc    interface{}
		update *types.UpdateAction
	}
	tests := []struct {
		name    string
		args    args
		expects []byte
		wantErr bool
	}{
		{
			name: "update action with json",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				[]byte(`{"id": 1, "name": "bar"}`),
				nil,
			},
			expects: []byte(`{"update":{"_index":"foo"}}
{"doc":{"id":1,"name":"bar"}}
`),
			wantErr: false,
		},
		{
			name: "update action with bad json",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				[]byte(`{"id": 1, "na`),
				nil,
			},
			expects: nil,
			wantErr: true,
		},
		{
			name: "update action with json.RawMessage",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				json.RawMessage(`{"id": 1, "name": "bar"}`),
				nil,
			},
			expects: []byte(`{"update":{"_index":"foo"}}
{"doc":{"id":1,"name":"bar"}}
`),
			wantErr: false,
		},
		{
			name: "update action with bad json.RawMessage",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				json.RawMessage(`{"id": 1, "na`),
				nil,
			},
			expects: nil,
			wantErr: true,
		},
		{
			name: "update action with struct",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{
					ID:   1,
					Name: "bar",
				},
				nil,
			},
			expects: []byte(`{"update":{"_index":"foo"}}
{"doc":{"id":1,"name":"bar"}}
`),
			wantErr: false,
		},
		{
			name: "update action with struct",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{
					ID:   1,
					Name: "bar",
				},
				&types.UpdateAction{},
			},
			expects: []byte(`{"update":{"_index":"foo"}}
{"doc":{"id":1,"name":"bar"}}
`),
			wantErr: false,
		},
		{
			name: "update action with struct and update.Doc",
			args: args{
				types.UpdateOperation{Index_: &indexName},
				struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{
					ID:   1,
					Name: "bar",
				},
				&types.UpdateAction{
					Doc: []byte(`{"id": 1, "name": "bar"}`),
				},
			},
			expects: []byte(`{"update":{"_index":"foo"}}
{"doc":{"id":1,"name":"bar"}}
`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.UpdateOp(tt.args.op, tt.args.doc, tt.args.update); (err != nil) != tt.wantErr {
				if r.buf.Len() == 0 {
					t.Errorf("buffer has not been reset")
				}
				t.Errorf("UpdateOp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(r.buf.Bytes(), tt.expects) {
				t.Errorf("serialized buffer invalid\nexpected: %s\ngot: %s", tt.expects, r.buf)
			}
		})
	}
}
