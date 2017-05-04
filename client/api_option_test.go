/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package client

import (
	"testing"

	"github.com/elastic/go-elasticsearch/api"
)

func TestOptionValidation(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}
	// TODO: maybe this should be an error type in some shared subpackage of api
	unsupportedOptionError := "unsupported option: WithIndex"
	_, err = c.Index("index", "good", nil, api.WithIndex("index again?"))
	if err.Error() != unsupportedOptionError {
		t.Fatalf("expected WithIndex() to return %q but got %q", unsupportedOptionError, err)
	}
	_, err = c.Index("index", "good", nil, api.WithHuman(true))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Index("index", "good", nil)
	if err != nil {
		t.Fatal(err)
	}
}
