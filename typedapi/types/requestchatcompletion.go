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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RequestChatCompletion type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L25-L97
type RequestChatCompletion struct {
	// MaxCompletionTokens The upper bound limit for the number of tokens that can be generated for a
	// completion request.
	MaxCompletionTokens *int64 `json:"max_completion_tokens,omitempty"`
	// Messages A list of objects representing the conversation.
	// Requests should generally only add new messages from the user (role `user`).
	// The other message roles (`assistant`, `system`, or `tool`) should generally
	// only be copied from the response to a previous completion request, such that
	// the messages array is built up throughout a conversation.
	Messages []Message `json:"messages"`
	// Model The ID of the model to use.
	Model *string `json:"model,omitempty"`
	// Stop A sequence of strings to control when the model should stop generating
	// additional tokens.
	Stop []string `json:"stop,omitempty"`
	// Temperature The sampling temperature to use.
	Temperature *float32 `json:"temperature,omitempty"`
	// ToolChoice Controls which tool is called by the model.
	// String representation: One of `auto`, `none`, or `requrired`. `auto` allows
	// the model to choose between calling tools and generating a message. `none`
	// causes the model to not call any tools. `required` forces the model to call
	// one or more tools.
	// Example (object representation):
	// ```
	//
	//	{
	//	  "tool_choice": {
	//	      "type": "function",
	//	      "function": {
	//	          "name": "get_current_weather"
	//	      }
	//	  }
	//	}
	//
	// ```
	ToolChoice CompletionToolType `json:"tool_choice,omitempty"`
	// Tools A list of tools that the model can call.
	// Example:
	// ```
	//
	//	{
	//	  "tools": [
	//	      {
	//	          "type": "function",
	//	          "function": {
	//	              "name": "get_price_of_item",
	//	              "description": "Get the current price of an item",
	//	              "parameters": {
	//	                  "type": "object",
	//	                  "properties": {
	//	                      "item": {
	//	                          "id": "12345"
	//	                      },
	//	                      "unit": {
	//	                          "type": "currency"
	//	                      }
	//	                  }
	//	              }
	//	          }
	//	      }
	//	  ]
	//	}
	//
	// ```
	Tools []CompletionTool `json:"tools,omitempty"`
	// TopP Nucleus sampling, an alternative to sampling with temperature.
	TopP *float32 `json:"top_p,omitempty"`
}

func (s *RequestChatCompletion) UnmarshalJSON(data []byte) error {

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

		case "max_completion_tokens":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxCompletionTokens", err)
				}
				s.MaxCompletionTokens = &value
			case float64:
				f := int64(v)
				s.MaxCompletionTokens = &f
			}

		case "messages":
			if err := dec.Decode(&s.Messages); err != nil {
				return fmt.Errorf("%s | %w", "Messages", err)
			}

		case "model":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Model", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Model = &o

		case "stop":
			if err := dec.Decode(&s.Stop); err != nil {
				return fmt.Errorf("%s | %w", "Stop", err)
			}

		case "temperature":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Temperature", err)
				}
				f := float32(value)
				s.Temperature = &f
			case float64:
				f := float32(v)
				s.Temperature = &f
			}

		case "tool_choice":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "ToolChoice", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		toolchoice_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "ToolChoice", err)
				}

				switch t {

				case "function", "type":
					o := NewCompletionToolChoice()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "ToolChoice", err)
					}
					s.ToolChoice = o
					break toolchoice_field

				}
			}
			if s.ToolChoice == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.ToolChoice); err != nil {
					return fmt.Errorf("%s | %w", "ToolChoice", err)
				}
			}

		case "tools":
			if err := dec.Decode(&s.Tools); err != nil {
				return fmt.Errorf("%s | %w", "Tools", err)
			}

		case "top_p":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopP", err)
				}
				f := float32(value)
				s.TopP = &f
			case float64:
				f := float32(v)
				s.TopP = &f
			}

		}
	}
	return nil
}

// NewRequestChatCompletion returns a RequestChatCompletion.
func NewRequestChatCompletion() *RequestChatCompletion {
	r := &RequestChatCompletion{}

	return r
}
