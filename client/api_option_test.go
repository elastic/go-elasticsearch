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
	"strings"
	"testing"

	"github.com/elastic/goelasticsearch/api"
)

func TestOptionValidation(t *testing.T) {
	c := New()
	// TODO: do something less hacky here
	postError := "Post http://localhost:9200: EOF"
	_, err := c.Index("index", "good", nil, api.WithIndex("index again?"))
	if err == nil || strings.HasSuffix(err.Error(), postError) {
		t.Fatalf("WithIndex was accepted as an argument to Index")
	}
	_, err = c.Index("index", "good", nil, api.WithHuman(true))
	if err != nil && !strings.HasSuffix(err.Error(), postError) {
		t.Fatal(err)
	}
	_, err = c.Index("index", "good", nil)
	if err != nil && !strings.HasSuffix(err.Error(), postError) {
		t.Fatal(err)
	}
}
