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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// PhraseSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/search/_types/suggester.ts#L356-L414
type PhraseSuggester struct {
	// Analyzer The analyzer to analyze the suggest text with.
	// Defaults to the search analyzer of the suggest field.
	Analyzer *string `json:"analyzer,omitempty"`
	// Collate Checks each suggestion against the specified query to prune suggestions for
	// which no matching docs exist in the index.
	Collate *PhraseSuggestCollate `json:"collate,omitempty"`
	// Confidence Defines a factor applied to the input phrases score, which is used as a
	// threshold for other suggest candidates.
	// Only candidates that score higher than the threshold will be included in the
	// result.
	Confidence *Float64 `json:"confidence,omitempty"`
	// DirectGenerator A list of candidate generators that produce a list of possible terms per term
	// in the given text.
	DirectGenerator []DirectGenerator `json:"direct_generator,omitempty"`
	// Field The field to fetch the candidate suggestions from.
	// Needs to be set globally or per suggestion.
	Field         string `json:"field"`
	ForceUnigrams *bool  `json:"force_unigrams,omitempty"`
	// GramSize Sets max size of the n-grams (shingles) in the field.
	// If the field doesn’t contain n-grams (shingles), this should be omitted or
	// set to `1`.
	// If the field uses a shingle filter, the `gram_size` is set to the
	// `max_shingle_size` if not explicitly set.
	GramSize *int `json:"gram_size,omitempty"`
	// Highlight Sets up suggestion highlighting.
	// If not provided, no highlighted field is returned.
	Highlight *PhraseSuggestHighlight `json:"highlight,omitempty"`
	// MaxErrors The maximum percentage of the terms considered to be misspellings in order to
	// form a correction.
	// This method accepts a float value in the range `[0..1)` as a fraction of the
	// actual query terms or a number `>=1` as an absolute number of query terms.
	MaxErrors *Float64 `json:"max_errors,omitempty"`
	// RealWordErrorLikelihood The likelihood of a term being misspelled even if the term exists in the
	// dictionary.
	RealWordErrorLikelihood *Float64 `json:"real_word_error_likelihood,omitempty"`
	// Separator The separator that is used to separate terms in the bigram field.
	// If not set, the whitespace character is used as a separator.
	Separator *string `json:"separator,omitempty"`
	// ShardSize Sets the maximum number of suggested terms to be retrieved from each
	// individual shard.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size The maximum corrections to be returned per suggest text token.
	Size *int `json:"size,omitempty"`
	// Smoothing The smoothing model used to balance weight between infrequent grams (grams
	// (shingles) are not existing in the index) and frequent grams (appear at least
	// once in the index).
	// The default model is Stupid Backoff.
	Smoothing *SmoothingModelContainer `json:"smoothing,omitempty"`
	// Text The text/query to provide suggestions for.
	Text       *string `json:"text,omitempty"`
	TokenLimit *int    `json:"token_limit,omitempty"`
}

func (s *PhraseSuggester) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "collate":
			if err := dec.Decode(&s.Collate); err != nil {
				return err
			}

		case "confidence":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Confidence = &f
			case float64:
				f := Float64(v)
				s.Confidence = &f
			}

		case "direct_generator":
			if err := dec.Decode(&s.DirectGenerator); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "force_unigrams":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ForceUnigrams = &value
			case bool:
				s.ForceUnigrams = &v
			}

		case "gram_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.GramSize = &value
			case float64:
				f := int(v)
				s.GramSize = &f
			}

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return err
			}

		case "max_errors":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.MaxErrors = &f
			case float64:
				f := Float64(v)
				s.MaxErrors = &f
			}

		case "real_word_error_likelihood":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.RealWordErrorLikelihood = &f
			case float64:
				f := Float64(v)
				s.RealWordErrorLikelihood = &f
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = &o

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "smoothing":
			if err := dec.Decode(&s.Smoothing); err != nil {
				return err
			}

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = &o

		case "token_limit":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TokenLimit = &value
			case float64:
				f := int(v)
				s.TokenLimit = &f
			}

		}
	}
	return nil
}

// NewPhraseSuggester returns a PhraseSuggester.
func NewPhraseSuggester() *PhraseSuggester {
	r := &PhraseSuggester{}

	return r
}
