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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// NodeInfoSettingsIngest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L86-L121
type NodeInfoSettingsIngest struct {
	Append          *NodeInfoIngestInfo `json:"append,omitempty"`
	Attachment      *NodeInfoIngestInfo `json:"attachment,omitempty"`
	Bytes           *NodeInfoIngestInfo `json:"bytes,omitempty"`
	Circle          *NodeInfoIngestInfo `json:"circle,omitempty"`
	Convert         *NodeInfoIngestInfo `json:"convert,omitempty"`
	Csv             *NodeInfoIngestInfo `json:"csv,omitempty"`
	Date            *NodeInfoIngestInfo `json:"date,omitempty"`
	DateIndexName   *NodeInfoIngestInfo `json:"date_index_name,omitempty"`
	Dissect         *NodeInfoIngestInfo `json:"dissect,omitempty"`
	DotExpander     *NodeInfoIngestInfo `json:"dot_expander,omitempty"`
	Drop            *NodeInfoIngestInfo `json:"drop,omitempty"`
	Enrich          *NodeInfoIngestInfo `json:"enrich,omitempty"`
	Fail            *NodeInfoIngestInfo `json:"fail,omitempty"`
	Foreach         *NodeInfoIngestInfo `json:"foreach,omitempty"`
	Geoip           *NodeInfoIngestInfo `json:"geoip,omitempty"`
	Grok            *NodeInfoIngestInfo `json:"grok,omitempty"`
	Gsub            *NodeInfoIngestInfo `json:"gsub,omitempty"`
	Inference       *NodeInfoIngestInfo `json:"inference,omitempty"`
	Join            *NodeInfoIngestInfo `json:"join,omitempty"`
	Json            *NodeInfoIngestInfo `json:"json,omitempty"`
	Kv              *NodeInfoIngestInfo `json:"kv,omitempty"`
	Lowercase       *NodeInfoIngestInfo `json:"lowercase,omitempty"`
	Pipeline        *NodeInfoIngestInfo `json:"pipeline,omitempty"`
	Remove          *NodeInfoIngestInfo `json:"remove,omitempty"`
	Rename          *NodeInfoIngestInfo `json:"rename,omitempty"`
	Script          *NodeInfoIngestInfo `json:"script,omitempty"`
	Set             *NodeInfoIngestInfo `json:"set,omitempty"`
	SetSecurityUser *NodeInfoIngestInfo `json:"set_security_user,omitempty"`
	Sort            *NodeInfoIngestInfo `json:"sort,omitempty"`
	Split           *NodeInfoIngestInfo `json:"split,omitempty"`
	Trim            *NodeInfoIngestInfo `json:"trim,omitempty"`
	Uppercase       *NodeInfoIngestInfo `json:"uppercase,omitempty"`
	Urldecode       *NodeInfoIngestInfo `json:"urldecode,omitempty"`
	UserAgent       *NodeInfoIngestInfo `json:"user_agent,omitempty"`
}

// NodeInfoSettingsIngestBuilder holds NodeInfoSettingsIngest struct and provides a builder API.
type NodeInfoSettingsIngestBuilder struct {
	v *NodeInfoSettingsIngest
}

// NewNodeInfoSettingsIngest provides a builder for the NodeInfoSettingsIngest struct.
func NewNodeInfoSettingsIngestBuilder() *NodeInfoSettingsIngestBuilder {
	r := NodeInfoSettingsIngestBuilder{
		&NodeInfoSettingsIngest{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsIngest struct
func (rb *NodeInfoSettingsIngestBuilder) Build() NodeInfoSettingsIngest {
	return *rb.v
}

func (rb *NodeInfoSettingsIngestBuilder) Append(append *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := append.Build()
	rb.v.Append = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Attachment(attachment *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := attachment.Build()
	rb.v.Attachment = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Bytes(bytes *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := bytes.Build()
	rb.v.Bytes = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Circle(circle *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := circle.Build()
	rb.v.Circle = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Convert(convert *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := convert.Build()
	rb.v.Convert = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Csv(csv *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := csv.Build()
	rb.v.Csv = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Date(date *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := date.Build()
	rb.v.Date = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) DateIndexName(dateindexname *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := dateindexname.Build()
	rb.v.DateIndexName = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Dissect(dissect *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := dissect.Build()
	rb.v.Dissect = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) DotExpander(dotexpander *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := dotexpander.Build()
	rb.v.DotExpander = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Drop(drop *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := drop.Build()
	rb.v.Drop = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Enrich(enrich *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := enrich.Build()
	rb.v.Enrich = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Fail(fail *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := fail.Build()
	rb.v.Fail = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Foreach(foreach *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := foreach.Build()
	rb.v.Foreach = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Geoip(geoip *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := geoip.Build()
	rb.v.Geoip = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Grok(grok *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := grok.Build()
	rb.v.Grok = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Gsub(gsub *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := gsub.Build()
	rb.v.Gsub = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Inference(inference *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := inference.Build()
	rb.v.Inference = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Join(join *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := join.Build()
	rb.v.Join = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Json(json *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := json.Build()
	rb.v.Json = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Kv(kv *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := kv.Build()
	rb.v.Kv = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Lowercase(lowercase *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := lowercase.Build()
	rb.v.Lowercase = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Pipeline(pipeline *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := pipeline.Build()
	rb.v.Pipeline = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Remove(remove *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := remove.Build()
	rb.v.Remove = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Rename(rename *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := rename.Build()
	rb.v.Rename = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Script(script *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Set(set *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := set.Build()
	rb.v.Set = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) SetSecurityUser(setsecurityuser *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := setsecurityuser.Build()
	rb.v.SetSecurityUser = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Sort(sort *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Split(split *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := split.Build()
	rb.v.Split = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Trim(trim *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := trim.Build()
	rb.v.Trim = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Uppercase(uppercase *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := uppercase.Build()
	rb.v.Uppercase = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) Urldecode(urldecode *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := urldecode.Build()
	rb.v.Urldecode = &v
	return rb
}

func (rb *NodeInfoSettingsIngestBuilder) UserAgent(useragent *NodeInfoIngestInfoBuilder) *NodeInfoSettingsIngestBuilder {
	v := useragent.Build()
	rb.v.UserAgent = &v
	return rb
}
