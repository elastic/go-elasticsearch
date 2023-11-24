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
)

// ScheduleContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Schedule.ts#L85-L96
type ScheduleContainer struct {
	Cron     *string         `json:"cron,omitempty"`
	Daily    *DailySchedule  `json:"daily,omitempty"`
	Hourly   *HourlySchedule `json:"hourly,omitempty"`
	Interval Duration        `json:"interval,omitempty"`
	Monthly  []TimeOfMonth   `json:"monthly,omitempty"`
	Weekly   []TimeOfWeek    `json:"weekly,omitempty"`
	Yearly   []TimeOfYear    `json:"yearly,omitempty"`
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
				return err
			}

		case "daily":
			if err := dec.Decode(&s.Daily); err != nil {
				return err
			}

		case "hourly":
			if err := dec.Decode(&s.Hourly); err != nil {
				return err
			}

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return err
			}

		case "monthly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfMonth()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Monthly = append(s.Monthly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Monthly); err != nil {
					return err
				}
			}

		case "weekly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfWeek()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Weekly = append(s.Weekly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Weekly); err != nil {
					return err
				}
			}

		case "yearly":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTimeOfYear()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Yearly = append(s.Yearly, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Yearly); err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// NewScheduleContainer returns a ScheduleContainer.
func NewScheduleContainer() *ScheduleContainer {
	r := &ScheduleContainer{}

	return r
}
