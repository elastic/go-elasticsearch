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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Message type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/inference/chat_completion_unified/UnifiedRequest.ts#L145-L165
type Message struct {
	// Content The content of the message.
	Content MessageContent `json:"content,omitempty"`
	// Role The role of the message author.
	Role string `json:"role"`
	// ToolCallId The tool call that this message is responding to.
	ToolCallId *string `json:"tool_call_id,omitempty"`
	// ToolCalls The tool calls generated by the model.
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

func (s *Message) UnmarshalJSON(data []byte) error {

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

		case "content":
			if err := dec.Decode(&s.Content); err != nil {
				return fmt.Errorf("%s | %w", "Content", err)
			}

		case "role":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Role", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Role = o

		case "tool_call_id":
			if err := dec.Decode(&s.ToolCallId); err != nil {
				return fmt.Errorf("%s | %w", "ToolCallId", err)
			}

		case "tool_calls":
			if err := dec.Decode(&s.ToolCalls); err != nil {
				return fmt.Errorf("%s | %w", "ToolCalls", err)
			}

		}
	}
	return nil
}

// NewMessage returns a Message.
func NewMessage() *Message {
	r := &Message{}

	return r
}

// true

type MessageVariant interface {
	MessageCaster() *Message
}

func (s *Message) MessageCaster() *Message {
	return s
}
