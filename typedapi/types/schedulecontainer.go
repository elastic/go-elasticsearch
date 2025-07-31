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

// ScheduleContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/_types/Schedule.ts#L80-L92
type ScheduleContainer struct {
	AdditionalScheduleContainerProperty map[string]json.RawMessage `json:"-"`
	Cron                                *string                    `json:"cron,omitempty"`
	Daily                               *DailySchedule             `json:"daily,omitempty"`
	Hourly                              *HourlySchedule            `json:"hourly,omitempty"`
	Interval                            Duration                   `json:"interval,omitempty"`
	Monthly                             []TimeOfMonth              `json:"monthly,omitempty"`
	Timezone                            *string                    `json:"timezone,omitempty"`
	Weekly                              []TimeOfWeek               `json:"weekly,omitempty"`
	Yearly                              []TimeOfYear               `json:"yearly,omitempty"`
}

func (s *ScheduleContainer) UnmarshalJSON(data []byte) error {

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

		case "cron":
			if err := dec.Decode(&s.Cron); err != nil {
				return fmt.Errorf("%s | %w", "Cron", err)
			}

		case "daily":
			if err := dec.Decode(&s.Daily); err != nil {
				return fmt.Errorf("%s | %w", "Daily", err)
			}

		case "hourly":
			if err := dec.Decode(&s.Hourly); err != nil {
				return fmt.Errorf("%s | %w", "Hourly", err)
			}

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return fmt.Errorf("%s | %w", "Interval", err)
			}

		case "monthly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfMonth()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Monthly", err)
				}

				s.Monthly = append(s.Monthly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Monthly); err != nil {
					return fmt.Errorf("%s | %w", "Monthly", err)
				}
			}

		case "timezone":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Timezone", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Timezone = &o

		case "weekly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfWeek()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Weekly", err)
				}

				s.Weekly = append(s.Weekly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Weekly); err != nil {
					return fmt.Errorf("%s | %w", "Weekly", err)
				}
			}

		case "yearly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfYear()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Yearly", err)
				}

				s.Yearly = append(s.Yearly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Yearly); err != nil {
					return fmt.Errorf("%s | %w", "Yearly", err)
				}
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalScheduleContainerProperty == nil {
					s.AdditionalScheduleContainerProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalScheduleContainerProperty", err)
				}
				s.AdditionalScheduleContainerProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ScheduleContainer) MarshalJSON() ([]byte, error) {
	type opt ScheduleContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalScheduleContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalScheduleContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewScheduleContainer returns a ScheduleContainer.
func NewScheduleContainer() *ScheduleContainer {
	r := &ScheduleContainer{
		AdditionalScheduleContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
