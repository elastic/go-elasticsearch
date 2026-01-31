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

package redact

import (
	"net/http"
	"strings"
)

// BasicAuth extracts basic auth credentials from the request and returns
// the username and a masked password. The password is masked to show only
// the first and last characters if it's 5+ characters long, otherwise
// it's fully masked. Returns empty strings and false if no auth is present.
func BasicAuth(r *http.Request) (username, maskedPassword string, ok bool) {
	username, password, ok := r.BasicAuth()
	if !ok {
		return "", "", false
	}
	return username, MaskPassword(password), true
}

// MaskPassword masks a password string, showing only the first and last
// characters if the password is 5+ characters long. Passwords shorter
// than 5 characters are fully masked.
func MaskPassword(password string) string {
	if len(password) < 5 {
		return strings.Repeat("*", len(password))
	}
	return string(password[0]) + strings.Repeat("*", len(password)-2) + string(password[len(password)-1])
}
