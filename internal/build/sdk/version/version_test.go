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

package version

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		versionStr string
	}
	tests := []struct {
		name    string
		args    args
		want    Version
		wantErr bool
	}{
		{
			name:    "9.0.1-SNAPSHOT",
			args:    args{versionStr: "9.0.1-SNAPSHOT"},
			want:    Version{9, 0, 1, "SNAPSHOT"},
			wantErr: false,
		},
		{
			name:    "9.0.1-rc1",
			args:    args{versionStr: "9.0.1-rc1"},
			want:    Version{9, 0, 1, "rc1"},
			wantErr: false,
		},
		{
			name:    "9.0.1",
			args:    args{versionStr: "9.0.1"},
			want:    Version{9, 0, 1, ""},
			wantErr: false,
		},
		{
			name:    "9.0.1.0",
			args:    args{versionStr: "9.0.1.0"},
			wantErr: true,
		},
		{
			name:    "9.0.1.0-SNAPSHOT",
			args:    args{versionStr: "9.0.1.0-SNAPSHOT"},
			wantErr: true,
		},
		{
			name:    "9.0-SNAPSHOT",
			args:    args{versionStr: "9.0-SNAPSHOT"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.versionStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
