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
	defaultPackage        = "api"
	templateFileName      = "method.tmpl"
	defaultPackageRepo    = "github.com/elastic/elasticsearch-go/api"
	transportPackageName  = "http"
	transportTypeName     = "Client"
	transportReceiverName = "c"
)

type apiPackages struct {
	packages map[string]*apiPackage
}

func newAPIPackages(methods []*method) *apiPackages {
	transportPackage := &apiPackage{
		PackageRepo:  transportPackageRepo,
		PackageName:  transportPackageName,
		TypeName:     transportTypeName,
		ReceiverName: transportReceiverName,
		FieldName:    transportFieldName,
	}
	a := &apiPackages{
		packages: map[string]*apiPackage{},
	}
	for _, m := range methods {
		if _, ok := a.packages[m.packageName]; ok {
			continue
		}
		p := newAPIPackage(m, transportPackage)
		a.packages[p.PackageName] = p
	}
	rootPackage := a.packages[defaultPackage]
	for _, p := range a.packages {
		if p.PackageName != rootPackage.PackageName {
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
