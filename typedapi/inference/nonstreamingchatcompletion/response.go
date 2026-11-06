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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package nonstreamingchatcompletion

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package nonstreamingchatcompletion
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/non_streaming_chat_completion/UnifiedResponse.ts#L22-L25
type Response struct {
	// Choices The list of completion choices the model generated for the input message.
	Choices []types.ChatCompletionChoice `json:"choices,omitempty"`
	// Id The unique identifier for the completion.
	Id string `json:"id"`
	// Model The model used to generate the completion.
	Model string `json:"model"`
	// Object The object type.
	Object string `json:"object"`
	// Usage The token usage statistics for the completion request. Omitted when usage
	// information is not available.
	Usage *types.ChatCompletionUsage `json:"usage,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
