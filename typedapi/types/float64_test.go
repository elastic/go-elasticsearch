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

package types

import (
	"math"
	"reflect"
	"testing"
)

func TestFloat64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		f       Float64
		want    string
		wantErr bool
	}{
		{
			name:    "standard",
			f:       1.0,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "standard",
			f:       1.1,
			want:    "1.1",
			wantErr: false,
		},
		{
			name:    "+Infinity",
			f:       Float64(math.Inf(+1)),
			want:    `"Infinity"`,
			wantErr: false,
		},
		{
			name:    "-Infinity",
			f:       Float64(math.Inf(-1)),
			want:    `"-Infinity"`,
			wantErr: false,
		},
		{
			name:    "NaN",
			f:       Float64(math.NaN()),
			want:    `"NaN"`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		f       Float64
		data    string
		wantErr bool
	}{
		{
			name:    "standard float",
			f:       1.0,
			data:    `1.0`,
			wantErr: false,
		},
		{
			name:    "standard float",
			f:       0.0,
			data:    `0.0`,
			wantErr: false,
		},
		{
			name:    "NaN",
			f:       Float64(math.NaN()),
			data:    "nan",
			wantErr: false,
		},
		{
			name:    "NaN",
			f:       Float64(math.NaN()),
			data:    `"NaN"`,
			wantErr: false,
		},
		{
			name:    "+Inf",
			f:       Float64(math.Inf(1)),
			data:    "Infinity",
			wantErr: false,
		},
		{
			name:    "+Inf",
			f:       Float64(math.Inf(1)),
			data:    `"Infinity"`,
			wantErr: false,
		},
		{
			name:    "-Inf",
			f:       Float64(math.Inf(-1)),
			data:    "-Infinity",
			wantErr: false,
		},
		{
			name:    "-Inf",
			f:       Float64(math.Inf(-1)),
			data:    `"-Infinity"`,
			wantErr: false,
		},
		{
			name:    "null",
			f:       0,
			data:    "null",
			wantErr: false,
		},
		{
			name:    "string",
			f:       0,
			data:    "not_a_float",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Float64
			err := f.UnmarshalJSON([]byte(tt.data))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if math.IsNaN(float64(tt.f)) {
				if !math.IsNaN(float64(f)) {
					t.Errorf("bad unmarshal, wanted: %v, got: %v", tt.f, f)
				}
			} else if tt.f != f {
				t.Errorf("bad unmarshal, wanted: %v, got: %v", tt.f, f)
			}
		})
	}
}
