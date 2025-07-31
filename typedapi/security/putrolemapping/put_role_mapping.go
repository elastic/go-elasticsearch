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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Create or update role mappings.
//
// Role mappings define which roles are assigned to each user.
// Each mapping has rules that identify users and a list of roles that are
// granted to those users.
// The role mapping APIs are generally the preferred way to manage role mappings
// rather than using role mapping files. The create or update role mappings API
// cannot update role mappings that are defined in role mapping files.
//
// NOTE: This API does not create roles. Rather, it maps users to existing
// roles.
// Roles can be created by using the create or update roles API or roles files.
//
// **Role templates**
//
// The most common use for role mappings is to create a mapping from a known
// value on the user to a fixed role name.
// For example, all users in the `cn=admin,dc=example,dc=com` LDAP group should
// be given the superuser role in Elasticsearch.
// The `roles` field is used for this purpose.
//
// For more complex needs, it is possible to use Mustache templates to
// dynamically determine the names of the roles that should be granted to the
// user.
// The `role_templates` field is used for this purpose.
//
// NOTE: To use role templates successfully, the relevant scripting feature must
// be enabled.
// Otherwise, all attempts to create a role mapping with role templates fail.
//
// All of the user fields that are available in the role mapping rules are also
// available in the role templates.
// Thus it is possible to assign a user to a role that reflects their username,
// their groups, or the name of the realm to which they authenticated.
//
// By default a template is evaluated to produce a single string that is the
// name of the role which should be assigned to the user.
// If the format of the template is set to "json" then the template is expected
// to produce a JSON string or an array of JSON strings for the role names.
package putrolemapping

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRoleMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutRoleMapping type alias for index.
type NewPutRoleMapping func(name string) *PutRoleMapping

// NewPutRoleMappingFunc returns a new instance of PutRoleMapping with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutRoleMappingFunc(tp elastictransport.Interface) NewPutRoleMapping {
	return func(name string) *PutRoleMapping {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update role mappings.
//
// Role mappings define which roles are assigned to each user.
// Each mapping has rules that identify users and a list of roles that are
// granted to those users.
// The role mapping APIs are generally the preferred way to manage role mappings
// rather than using role mapping files. The create or update role mappings API
// cannot update role mappings that are defined in role mapping files.
//
// NOTE: This API does not create roles. Rather, it maps users to existing
// roles.
// Roles can be created by using the create or update roles API or roles files.
//
// **Role templates**
//
// The most common use for role mappings is to create a mapping from a known
// value on the user to a fixed role name.
// For example, all users in the `cn=admin,dc=example,dc=com` LDAP group should
// be given the superuser role in Elasticsearch.
// The `roles` field is used for this purpose.
//
// For more complex needs, it is possible to use Mustache templates to
// dynamically determine the names of the roles that should be granted to the
// user.
// The `role_templates` field is used for this purpose.
//
// NOTE: To use role templates successfully, the relevant scripting feature must
// be enabled.
// Otherwise, all attempts to create a role mapping with role templates fail.
//
// All of the user fields that are available in the role mapping rules are also
// available in the role templates.
// Thus it is possible to assign a user to a role that reflects their username,
// their groups, or the name of the realm to which they authenticated.
//
// By default a template is evaluated to produce a single string that is the
// name of the role which should be assigned to the user.
// If the format of the template is set to "json" then the template is expected
// to produce a JSON string or an array of JSON strings for the role names.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html
func New(tp elastictransport.Interface) *PutRoleMapping {
	r := &PutRoleMapping{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutRoleMapping) Raw(raw io.Reader) *PutRoleMapping {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutRoleMapping) Request(req *Request) *PutRoleMapping {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutRoleMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutRoleMapping: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("role_mapping")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutRoleMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.put_role_mapping")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "security.put_role_mapping")
		if reader := instrument.RecordRequestBody(ctx, "security.put_role_mapping", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.put_role_mapping")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutRoleMapping query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putrolemapping.Response
func (r PutRoleMapping) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.put_role_mapping")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the PutRoleMapping headers map.
func (r *PutRoleMapping) Header(key, value string) *PutRoleMapping {
	r.headers.Set(key, value)

	return r
}

// Name The distinct name that identifies the role mapping.
// The name is used solely as an identifier to facilitate interaction via the
// API; it does not affect the behavior of the mapping in any way.
// API Name: name
func (r *PutRoleMapping) _name(name string) *PutRoleMapping {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Refresh If `true` (the default) then refresh the affected shards to make this
// operation visible to search, if `wait_for` then wait for a refresh to make
// this operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *PutRoleMapping) Refresh(refresh refresh.Refresh) *PutRoleMapping {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutRoleMapping) ErrorTrace(errortrace bool) *PutRoleMapping {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutRoleMapping) FilterPath(filterpaths ...string) *PutRoleMapping {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *PutRoleMapping) Human(human bool) *PutRoleMapping {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutRoleMapping) Pretty(pretty bool) *PutRoleMapping {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Enabled Mappings that have `enabled` set to `false` are ignored when role mapping is
// performed.
// API name: enabled
func (r *PutRoleMapping) Enabled(enabled bool) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Enabled = &enabled

	return r
}

// Metadata Additional metadata that helps define which roles are assigned to each user.
// Within the metadata object, keys beginning with `_` are reserved for system
// usage.
// API name: metadata
func (r *PutRoleMapping) Metadata(metadata types.Metadata) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// RoleTemplates A list of Mustache templates that will be evaluated to determine the roles
// names that should granted to the users that match the role mapping rules.
// Exactly one of `roles` or `role_templates` must be specified.
// API name: role_templates
func (r *PutRoleMapping) RoleTemplates(roletemplates ...types.RoleTemplate) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RoleTemplates = roletemplates

	return r
}

// Roles A list of role names that are granted to the users that match the role
// mapping rules.
// Exactly one of `roles` or `role_templates` must be specified.
// API name: roles
func (r *PutRoleMapping) Roles(roles ...string) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Roles = roles

	return r
}

// Rules The rules that determine which users should be matched by the mapping.
// A rule is a logical condition that is expressed by using a JSON DSL.
// API name: rules
func (r *PutRoleMapping) Rules(rules *types.RoleMappingRule) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Rules = rules

	return r
}

// API name: run_as
func (r *PutRoleMapping) RunAs(runas ...string) *PutRoleMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RunAs = runas

	return r
}
