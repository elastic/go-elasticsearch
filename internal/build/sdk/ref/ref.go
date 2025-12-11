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
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/internal/build/sdk/version"
)

type Ref struct {
	branch  string
	version *version.Version
}

func (r Ref) IsEmpty() bool {
	return r.version == nil && r.branch == ""
}

func (r Ref) IsPreRelease() bool {
	return (r.version != nil && r.version.IsPreRelease()) || r.branch != ""
}

// Parse parses a reference string into a Ref. The reference can be a semantic
// version (e.g., "8.15.0", "9.0.0-SNAPSHOT") or a branch name (e.g., "main", "master").
// An optional rename function can be provided to transform branch names during parsing.
func Parse(reference string, rename ...func(string) string) (Ref, error) {
	if reference == "" {
		return Ref{}, errors.New("empty reference")
	}

	// Try to parse as a version first
	if ver, err := version.Parse(reference); err == nil {
		return Ref{version: &ver}, nil
	}

	// Otherwise, treat as a branch name
	branch := reference
	if len(rename) > 0 && rename[0] != nil {
		branch = rename[0](reference)
	}
	return Ref{branch: branch}, nil
}

func (r Ref) TargetBranch() string {
	if r.branch != "" {
		return r.branch
	}
	return fmt.Sprintf("%d.%d", r.version.Major(), r.version.Minor())
}

func (r Ref) String() string {
	if r.version != nil {
		return r.version.String()
	}
	return r.branch
}

// BaseVersion returns the version string without any prerelease suffix.
// For branch refs, it returns an empty string.
func (r Ref) BaseVersion() string {
	if r.version != nil {
		return r.version.WithPreRelease("").String()
	}
	return ""
}
