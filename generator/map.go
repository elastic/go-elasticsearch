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

import (
	"fmt"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"packageName":  packageName,
	"methodName":   methodName,
	"typeName":     typeName,
	"receiverName": receiverName,
	"httpMethod":   httpMethod,
}

// packageAndMethod is a helper for other methods obtaining the package and/or method from the <api>.<method> format
func packageAndMethod(spec map[string]interface{}) (string, string, error) {
	api, ok := spec["method"].(string)
	if !ok {
		return "", "", fmt.Errorf("Unexpected type for method: %T (expected string)", spec["method"])
	}
	apiParts := strings.Split(api, ".")
	packageName := defaultPackage
	var methodName string
	switch len(apiParts) {
	case 1:
		methodName = apiParts[0]
	case 2:
		packageName = apiParts[0]
		methodName = apiParts[1]
	default:
		return "", "", fmt.Errorf("Unexpected API name format: %s", api)
	}
	return packageName, methodName, nil
}

func packageName(spec map[string]interface{}) (string, error) {
	p, _, err := packageAndMethod(spec)
	return p, err
}

func methodName(spec map[string]interface{}) (string, error) {
	_, m, err := packageAndMethod(spec)
	return toCamel(m), err
}

func typeName(spec map[string]interface{}) (string, error) {
	p, _, err := packageAndMethod(spec)
	if err != nil {
		return "", err
	}
	return toCamel(p), nil
}

func receiverName(spec map[string]interface{}) (string, error) {
	t, err := typeName(spec)
	if err != nil {
		return "", err
	}
	return strings.ToLower(string(t[0])), nil
}

func httpMethod(spec map[string]interface{}) (string, error) {
	m, ok := spec["spec"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Unexpected type for method: %T (expected map[string]interface {})", spec["spec"])
	}
	methods, ok := m["methods"].([]interface{})
	if !ok {
		return "", fmt.Errorf("Unexpected type for methods: %T (expected []interface{})", m["methods"])
	}
	for _, m := range methods {
		methodName, ok := m.(string)
		if !ok {
			return "", fmt.Errorf("Unexpected type for method: %T (expected string)", m)
		}
		return methodName, nil
	}
	return "", fmt.Errorf("No HTTP methods in %q", spec)
}
