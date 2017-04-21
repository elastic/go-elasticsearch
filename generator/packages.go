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

package generator

const (
	defaultPackage     = "client"
	templateFileName   = "method.tmpl"
	defaultPackageRepo = "github.com/elastic/elasticsearch-go/client"
)

type apiPackages struct {
	packages map[string]*apiPackage
}

func newAPIPackages(methods []*method) *apiPackages {
	a := &apiPackages{
		packages: map[string]*apiPackage{},
	}
	for _, m := range methods {
		if p, ok := a.packages[m.PackageName]; ok {
			p.addMethod(m)
		} else {
			a.packages[m.PackageName] = newAPIPackage(m)
		}
	}
	rootPackage := a.packages[defaultPackage]
	for _, p := range a.packages {
		if p.Methods[0].PackageName != rootPackage.Methods[0].PackageName {
			rootPackage.addSubpackage(p)
		}
	}
	return a
}

func (a *apiPackages) generate(templatesDir, outputDir string) error {
	for _, p := range a.packages {
		if err := p.generate(templatesDir, outputDir); err != nil {
			return err
		}
	}
	return nil
}
