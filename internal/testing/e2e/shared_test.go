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

//go:build integration
// +build integration

package e2e_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/elastic/go-elasticsearch/v9"

	"testing/containertest"
)

// sharedCfg is the Elasticsearch client config for the shared secure container.
var sharedCfg elasticsearch.Config

// sharedInsecureCfg is the Elasticsearch client config for the shared insecure
// container (xpack.security.enabled=false).
var sharedInsecureCfg elasticsearch.Config

// sharedInsecureAddr is the raw address of the insecure container, used by
// tests that need to construct transports or parse the URL directly.
var sharedInsecureAddr string

func TestMain(m *testing.M) {
	var (
		secureSrv   *containertest.ElasticsearchService
		insecureSrv *containertest.ElasticsearchService

		secureErr   error
		insecureErr error
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		secureSrv, secureErr = containertest.NewElasticsearchService(containertest.ElasticStackImage)
	}()

	go func() {
		defer wg.Done()
		insecureSrv, insecureErr = containertest.NewElasticsearchService(
			containertest.ElasticStackImage,
			containertest.WithEnv(map[string]string{
				"xpack.security.enabled": "false",
			}),
		)
	}()

	wg.Wait()

	if secureErr != nil {
		log.Fatalf("Error setting up shared secure Elasticsearch container: %s", secureErr)
	}
	if insecureErr != nil {
		log.Fatalf("Error setting up shared insecure Elasticsearch container: %s", insecureErr)
	}

	sharedCfg = secureSrv.ESConfig()
	insecureCfg := insecureSrv.ESConfig()
	sharedInsecureCfg = insecureCfg
	sharedInsecureAddr = insecureCfg.Addresses[0]

	code := m.Run()

	errs := make([]error, 0, 2)
	if err := secureSrv.Terminate(context.Background()); err != nil {
		errs = append(errs, fmt.Errorf("secure container: %w", err))
	}
	if err := insecureSrv.Terminate(context.Background()); err != nil {
		errs = append(errs, fmt.Errorf("insecure container: %w", err))
	}
	for _, err := range errs {
		log.Printf("Error terminating container: %s", err)
	}

	os.Exit(code)
}
