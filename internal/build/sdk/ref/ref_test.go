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

package ref

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		reference string
		rename    func(string) string
		wantStr   string
		wantErr   bool
	}{
		{
			name:      "valid semantic version",
			reference: "8.15.0",
			wantStr:   "8.15.0",
			wantErr:   false,
		},
		{
			name:      "valid semantic version with snapshot",
			reference: "9.0.0-SNAPSHOT",
			wantStr:   "9.0.0-SNAPSHOT",
			wantErr:   false,
		},
		{
			name:      "valid semantic version with rc",
			reference: "8.15.0-rc1",
			wantStr:   "8.15.0-rc1",
			wantErr:   false,
		},
		{
			name:      "branch name main",
			reference: "main",
			wantStr:   "main",
			wantErr:   false,
		},
		{
			name:      "branch name master",
			reference: "master",
			wantStr:   "master",
			wantErr:   false,
		},
		{
			name:      "branch name with rename function",
			reference: "main",
			rename:    func(s string) string { return "renamed-" + s },
			wantStr:   "renamed-main",
			wantErr:   false,
		},
		{
			name:      "version is not affected by rename function",
			reference: "8.15.0",
			rename:    func(s string) string { return "renamed-" + s },
			wantStr:   "8.15.0",
			wantErr:   false,
		},
		{
			name:      "feature branch",
			reference: "feature/new-api",
			wantStr:   "feature/new-api",
			wantErr:   false,
		},
		{
			name:      "empty reference",
			reference: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Ref
			var err error
			if tt.rename != nil {
				got, err = Parse(tt.reference, tt.rename)
			} else {
				got, err = Parse(tt.reference)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got.String() != tt.wantStr {
				t.Errorf("Parse() got = %v, want %v", got.String(), tt.wantStr)
			}
		})
	}
}

func TestRef_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		ref  Ref
		want bool
	}{
		{
			name: "empty ref",
			ref:  Ref{},
			want: true,
		},
		{
			name: "ref with branch",
			ref:  mustParse("main"),
			want: false,
		},
		{
			name: "ref with version",
			ref:  mustParse("8.15.0"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.IsEmpty(); got != tt.want {
				t.Errorf("Ref.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_IsPreRelease(t *testing.T) {
	tests := []struct {
		name string
		ref  Ref
		want bool
	}{
		{
			name: "stable version is not pre-release",
			ref:  mustParse("8.15.0"),
			want: false,
		},
		{
			name: "snapshot version is pre-release",
			ref:  mustParse("9.0.0-SNAPSHOT"),
			want: true,
		},
		{
			name: "rc version is pre-release",
			ref:  mustParse("8.15.0-rc1"),
			want: true,
		},
		{
			name: "alpha version is pre-release",
			ref:  mustParse("8.15.0-alpha1"),
			want: true,
		},
		{
			name: "branch is pre-release",
			ref:  mustParse("main"),
			want: true,
		},
		{
			name: "feature branch is pre-release",
			ref:  mustParse("feature/new-api"),
			want: true,
		},
		{
			name: "empty ref is not pre-release",
			ref:  Ref{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.IsPreRelease(); got != tt.want {
				t.Errorf("Ref.IsPreRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_TargetBranch(t *testing.T) {
	tests := []struct {
		name string
		ref  Ref
		want string
	}{
		{
			name: "version returns major.minor",
			ref:  mustParse("8.15.0"),
			want: "8.15",
		},
		{
			name: "snapshot version returns major.minor",
			ref:  mustParse("9.0.0-SNAPSHOT"),
			want: "9.0",
		},
		{
			name: "branch returns branch name",
			ref:  mustParse("main"),
			want: "main",
		},
		{
			name: "feature branch returns branch name",
			ref:  mustParse("feature/new-api"),
			want: "feature/new-api",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.TargetBranch(); got != tt.want {
				t.Errorf("Ref.TargetBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_String(t *testing.T) {
	tests := []struct {
		name string
		ref  Ref
		want string
	}{
		{
			name: "version returns version string",
			ref:  mustParse("8.15.0"),
			want: "8.15.0",
		},
		{
			name: "snapshot version returns full version string",
			ref:  mustParse("9.0.0-SNAPSHOT"),
			want: "9.0.0-SNAPSHOT",
		},
		{
			name: "branch returns branch name",
			ref:  mustParse("main"),
			want: "main",
		},
		{
			name: "empty ref returns empty string",
			ref:  Ref{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.String(); got != tt.want {
				t.Errorf("Ref.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_BaseVersion(t *testing.T) {
	tests := []struct {
		name string
		ref  Ref
		want string
	}{
		{
			name: "stable version returns same version",
			ref:  mustParse("8.15.0"),
			want: "8.15.0",
		},
		{
			name: "snapshot version returns version without prerelease",
			ref:  mustParse("9.0.0-SNAPSHOT"),
			want: "9.0.0",
		},
		{
			name: "rc version returns version without prerelease",
			ref:  mustParse("8.15.0-rc1"),
			want: "8.15.0",
		},
		{
			name: "alpha version returns version without prerelease",
			ref:  mustParse("8.15.0-alpha1"),
			want: "8.15.0",
		},
		{
			name: "branch returns empty string",
			ref:  mustParse("main"),
			want: "",
		},
		{
			name: "empty ref returns empty string",
			ref:  Ref{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.BaseVersion(); got != tt.want {
				t.Errorf("Ref.BaseVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse_WithRenameFunction(t *testing.T) {
	// Test that the rename function is applied correctly to branch names
	uppercaseRename := func(s string) string {
		return strings.ToUpper(s)
	}

	tests := []struct {
		name      string
		reference string
		rename    func(string) string
		wantStr   string
	}{
		{
			name:      "rename main to MAIN",
			reference: "main",
			rename:    uppercaseRename,
			wantStr:   "MAIN",
		},
		{
			name:      "nil rename function does not modify branch",
			reference: "main",
			rename:    nil,
			wantStr:   "main",
		},
		{
			name:      "version is not renamed",
			reference: "8.15.0",
			rename:    uppercaseRename,
			wantStr:   "8.15.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.reference, tt.rename)
			if err != nil {
				t.Errorf("Parse() unexpected error = %v", err)
				return
			}
			if got.String() != tt.wantStr {
				t.Errorf("Parse() got = %v, want %v", got.String(), tt.wantStr)
			}
		})
	}
}

func mustParse(reference string) Ref {
	ref, err := Parse(reference)
	if err != nil {
		panic(err)
	}
	return ref
}
