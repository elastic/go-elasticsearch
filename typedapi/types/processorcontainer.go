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

// ProcessorContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L28-L67
type ProcessorContainer struct {
	Append          *AppendProcessor          `json:"append,omitempty"`
	Attachment      *AttachmentProcessor      `json:"attachment,omitempty"`
	Bytes           *BytesProcessor           `json:"bytes,omitempty"`
	Circle          *CircleProcessor          `json:"circle,omitempty"`
	Convert         *ConvertProcessor         `json:"convert,omitempty"`
	Csv             *CsvProcessor             `json:"csv,omitempty"`
	Date            *DateProcessor            `json:"date,omitempty"`
	DateIndexName   *DateIndexNameProcessor   `json:"date_index_name,omitempty"`
	Dissect         *DissectProcessor         `json:"dissect,omitempty"`
	DotExpander     *DotExpanderProcessor     `json:"dot_expander,omitempty"`
	Drop            *DropProcessor            `json:"drop,omitempty"`
	Enrich          *EnrichProcessor          `json:"enrich,omitempty"`
	Fail            *FailProcessor            `json:"fail,omitempty"`
	Foreach         *ForeachProcessor         `json:"foreach,omitempty"`
	Geoip           *GeoIpProcessor           `json:"geoip,omitempty"`
	Grok            *GrokProcessor            `json:"grok,omitempty"`
	Gsub            *GsubProcessor            `json:"gsub,omitempty"`
	Inference       *InferenceProcessor       `json:"inference,omitempty"`
	Join            *JoinProcessor            `json:"join,omitempty"`
	Json            *JsonProcessor            `json:"json,omitempty"`
	Kv              *KeyValueProcessor        `json:"kv,omitempty"`
	Lowercase       *LowercaseProcessor       `json:"lowercase,omitempty"`
	Pipeline        *PipelineProcessor        `json:"pipeline,omitempty"`
	Remove          *RemoveProcessor          `json:"remove,omitempty"`
	Rename          *RenameProcessor          `json:"rename,omitempty"`
	Script          *Script                   `json:"script,omitempty"`
	Set             *SetProcessor             `json:"set,omitempty"`
	SetSecurityUser *SetSecurityUserProcessor `json:"set_security_user,omitempty"`
	Sort            *SortProcessor            `json:"sort,omitempty"`
	Split           *SplitProcessor           `json:"split,omitempty"`
	Trim            *TrimProcessor            `json:"trim,omitempty"`
	Uppercase       *UppercaseProcessor       `json:"uppercase,omitempty"`
	Urldecode       *UrlDecodeProcessor       `json:"urldecode,omitempty"`
	UserAgent       *UserAgentProcessor       `json:"user_agent,omitempty"`
}

// ProcessorContainerBuilder holds ProcessorContainer struct and provides a builder API.
type ProcessorContainerBuilder struct {
	v *ProcessorContainer
}

// NewProcessorContainer provides a builder for the ProcessorContainer struct.
func NewProcessorContainerBuilder() *ProcessorContainerBuilder {
	r := ProcessorContainerBuilder{
		&ProcessorContainer{},
	}

	return &r
}

// Build finalize the chain and returns the ProcessorContainer struct
func (rb *ProcessorContainerBuilder) Build() ProcessorContainer {
	return *rb.v
}

func (rb *ProcessorContainerBuilder) Append(append *AppendProcessorBuilder) *ProcessorContainerBuilder {
	v := append.Build()
	rb.v.Append = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Attachment(attachment *AttachmentProcessorBuilder) *ProcessorContainerBuilder {
	v := attachment.Build()
	rb.v.Attachment = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Bytes(bytes *BytesProcessorBuilder) *ProcessorContainerBuilder {
	v := bytes.Build()
	rb.v.Bytes = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Circle(circle *CircleProcessorBuilder) *ProcessorContainerBuilder {
	v := circle.Build()
	rb.v.Circle = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Convert(convert *ConvertProcessorBuilder) *ProcessorContainerBuilder {
	v := convert.Build()
	rb.v.Convert = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Csv(csv *CsvProcessorBuilder) *ProcessorContainerBuilder {
	v := csv.Build()
	rb.v.Csv = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Date(date *DateProcessorBuilder) *ProcessorContainerBuilder {
	v := date.Build()
	rb.v.Date = &v
	return rb
}

func (rb *ProcessorContainerBuilder) DateIndexName(dateindexname *DateIndexNameProcessorBuilder) *ProcessorContainerBuilder {
	v := dateindexname.Build()
	rb.v.DateIndexName = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Dissect(dissect *DissectProcessorBuilder) *ProcessorContainerBuilder {
	v := dissect.Build()
	rb.v.Dissect = &v
	return rb
}

func (rb *ProcessorContainerBuilder) DotExpander(dotexpander *DotExpanderProcessorBuilder) *ProcessorContainerBuilder {
	v := dotexpander.Build()
	rb.v.DotExpander = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Drop(drop *DropProcessorBuilder) *ProcessorContainerBuilder {
	v := drop.Build()
	rb.v.Drop = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Enrich(enrich *EnrichProcessorBuilder) *ProcessorContainerBuilder {
	v := enrich.Build()
	rb.v.Enrich = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Fail(fail *FailProcessorBuilder) *ProcessorContainerBuilder {
	v := fail.Build()
	rb.v.Fail = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Foreach(foreach *ForeachProcessorBuilder) *ProcessorContainerBuilder {
	v := foreach.Build()
	rb.v.Foreach = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Geoip(geoip *GeoIpProcessorBuilder) *ProcessorContainerBuilder {
	v := geoip.Build()
	rb.v.Geoip = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Grok(grok *GrokProcessorBuilder) *ProcessorContainerBuilder {
	v := grok.Build()
	rb.v.Grok = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Gsub(gsub *GsubProcessorBuilder) *ProcessorContainerBuilder {
	v := gsub.Build()
	rb.v.Gsub = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Inference(inference *InferenceProcessorBuilder) *ProcessorContainerBuilder {
	v := inference.Build()
	rb.v.Inference = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Join(join *JoinProcessorBuilder) *ProcessorContainerBuilder {
	v := join.Build()
	rb.v.Join = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Json(json *JsonProcessorBuilder) *ProcessorContainerBuilder {
	v := json.Build()
	rb.v.Json = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Kv(kv *KeyValueProcessorBuilder) *ProcessorContainerBuilder {
	v := kv.Build()
	rb.v.Kv = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Lowercase(lowercase *LowercaseProcessorBuilder) *ProcessorContainerBuilder {
	v := lowercase.Build()
	rb.v.Lowercase = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Pipeline(pipeline *PipelineProcessorBuilder) *ProcessorContainerBuilder {
	v := pipeline.Build()
	rb.v.Pipeline = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Remove(remove *RemoveProcessorBuilder) *ProcessorContainerBuilder {
	v := remove.Build()
	rb.v.Remove = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Rename(rename *RenameProcessorBuilder) *ProcessorContainerBuilder {
	v := rename.Build()
	rb.v.Rename = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Script(script *ScriptBuilder) *ProcessorContainerBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Set(set *SetProcessorBuilder) *ProcessorContainerBuilder {
	v := set.Build()
	rb.v.Set = &v
	return rb
}

func (rb *ProcessorContainerBuilder) SetSecurityUser(setsecurityuser *SetSecurityUserProcessorBuilder) *ProcessorContainerBuilder {
	v := setsecurityuser.Build()
	rb.v.SetSecurityUser = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Sort(sort *SortProcessorBuilder) *ProcessorContainerBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Split(split *SplitProcessorBuilder) *ProcessorContainerBuilder {
	v := split.Build()
	rb.v.Split = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Trim(trim *TrimProcessorBuilder) *ProcessorContainerBuilder {
	v := trim.Build()
	rb.v.Trim = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Uppercase(uppercase *UppercaseProcessorBuilder) *ProcessorContainerBuilder {
	v := uppercase.Build()
	rb.v.Uppercase = &v
	return rb
}

func (rb *ProcessorContainerBuilder) Urldecode(urldecode *UrlDecodeProcessorBuilder) *ProcessorContainerBuilder {
	v := urldecode.Build()
	rb.v.Urldecode = &v
	return rb
}

func (rb *ProcessorContainerBuilder) UserAgent(useragent *UserAgentProcessorBuilder) *ProcessorContainerBuilder {
	v := useragent.Build()
	rb.v.UserAgent = &v
	return rb
}
