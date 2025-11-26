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
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	versionRexep = regexp.MustCompile(`^(?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)(?:-(?P<prerelease>[0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?$`)
)

// Version represents a semantic version.
type Version struct {
	major, minor, patch int
	preRelease          string
}

// Major returns the major version.
func (v Version) Major() int {
	return v.major
}

// Minor returns the minor version.
func (v Version) Minor() int {
	return v.minor
}

// Patch returns the patch version.
func (v Version) Patch() int {
	return v.patch
}

// PreRelease returns the pre-release version, if any.
func (v Version) PreRelease() string {
	return v.preRelease
}

// IsPreRelease returns true if the version is a pre-release version. e.g. 7.10.0-alpha1, or 7.10.0-SNAPSHOT
func (v Version) IsPreRelease() bool {
	return v.PreRelease() != ""
}

// WithMajor returns a new Version with the given major version.
func (v Version) WithMajor(major int) Version {
	return Version{major: major, minor: v.minor, patch: v.patch, preRelease: v.preRelease}
}

// WithMinor returns a new Version with the given minor version.
func (v Version) WithMinor(minor int) Version {
	return Version{major: v.major, minor: minor, patch: v.patch, preRelease: v.preRelease}
}

// WithPatch returns a new Version with the given patch version.
func (v Version) WithPatch(patch int) Version {
	return Version{major: v.major, minor: v.minor, patch: patch, preRelease: v.preRelease}
}

// WithPreRelease returns a new Version with the given pre-release version.
func (v Version) WithPreRelease(preRelease string) Version {
	return Version{major: v.major, minor: v.minor, patch: v.patch, preRelease: preRelease}
}

// String returns the version as a string.
func (v Version) String() string {
	base := strings.Join(
		[]string{
			fmt.Sprint(v.major),
			fmt.Sprint(v.minor),
			fmt.Sprint(v.patch),
		}, ".")

	if v.IsPreRelease() {
		return fmt.Sprintf("%s-%s", base, v.preRelease)
	}
	return base
}

// parseIntField extracts and parses an integer field from the regex match results.
func parseIntField(result map[string]string, field string) (int, error) {
	if val, ok := result[field]; ok && val != "" {
		return strconv.Atoi(val)
	}
	return 0, nil
}

// Parse parses a version string into a Version struct.
func Parse(versionStr string) (Version, error) {
	matches := versionRexep.FindStringSubmatch(versionStr)
	if matches == nil {
		return Version{}, errors.New("invalid version format")
	}

	groupNames := versionRexep.SubexpNames()
	result := make(map[string]string)
	for i, name := range groupNames {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	major, err := parseIntField(result, "major")
	if err != nil {
		return Version{}, fmt.Errorf("invalid major version: %w", err)
	}

	minor, err := parseIntField(result, "minor")
	if err != nil {
		return Version{}, fmt.Errorf("invalid minor version: %w", err)
	}

	patch, err := parseIntField(result, "patch")
	if err != nil {
		return Version{}, fmt.Errorf("invalid patch version: %w", err)
	}

	return Version{
		major:      major,
		minor:      minor,
		patch:      patch,
		preRelease: result["prerelease"],
	}, nil
}

// MustParse parses a version string into a Version struct.
// It panics if the version string is invalid.
func MustParse(version string) Version {
	v, err := Parse(version)
	if err != nil {
		panic(err)
	}
	return v
}
