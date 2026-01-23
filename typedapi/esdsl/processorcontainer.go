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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _processorContainer struct {
	v *types.ProcessorContainer
}

func NewProcessorContainer() *_processorContainer {
	return &_processorContainer{v: types.NewProcessorContainer()}
}

// AdditionalProcessorContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_processorContainer) AdditionalProcessorContainerProperty(key string, value json.RawMessage) *_processorContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalProcessorContainerProperty = tmp
	return s
}

func (s *_processorContainer) Append(append types.AppendProcessorVariant) *_processorContainer {

	s.v.Append = append.AppendProcessorCaster()

	return s
}

func (s *_processorContainer) Attachment(attachment types.AttachmentProcessorVariant) *_processorContainer {

	s.v.Attachment = attachment.AttachmentProcessorCaster()

	return s
}

func (s *_processorContainer) Bytes(bytes types.BytesProcessorVariant) *_processorContainer {

	s.v.Bytes = bytes.BytesProcessorCaster()

	return s
}

func (s *_processorContainer) Cef(cef types.CefProcessorVariant) *_processorContainer {

	s.v.Cef = cef.CefProcessorCaster()

	return s
}

func (s *_processorContainer) Circle(circle types.CircleProcessorVariant) *_processorContainer {

	s.v.Circle = circle.CircleProcessorCaster()

	return s
}

func (s *_processorContainer) CommunityId(communityid types.CommunityIDProcessorVariant) *_processorContainer {

	s.v.CommunityId = communityid.CommunityIDProcessorCaster()

	return s
}

func (s *_processorContainer) Convert(convert types.ConvertProcessorVariant) *_processorContainer {

	s.v.Convert = convert.ConvertProcessorCaster()

	return s
}

func (s *_processorContainer) Csv(csv types.CsvProcessorVariant) *_processorContainer {

	s.v.Csv = csv.CsvProcessorCaster()

	return s
}

func (s *_processorContainer) Date(date types.DateProcessorVariant) *_processorContainer {

	s.v.Date = date.DateProcessorCaster()

	return s
}

func (s *_processorContainer) DateIndexName(dateindexname types.DateIndexNameProcessorVariant) *_processorContainer {

	s.v.DateIndexName = dateindexname.DateIndexNameProcessorCaster()

	return s
}

func (s *_processorContainer) Dissect(dissect types.DissectProcessorVariant) *_processorContainer {

	s.v.Dissect = dissect.DissectProcessorCaster()

	return s
}

func (s *_processorContainer) DotExpander(dotexpander types.DotExpanderProcessorVariant) *_processorContainer {

	s.v.DotExpander = dotexpander.DotExpanderProcessorCaster()

	return s
}

func (s *_processorContainer) Drop(drop types.DropProcessorVariant) *_processorContainer {

	s.v.Drop = drop.DropProcessorCaster()

	return s
}

func (s *_processorContainer) Enrich(enrich types.EnrichProcessorVariant) *_processorContainer {

	s.v.Enrich = enrich.EnrichProcessorCaster()

	return s
}

func (s *_processorContainer) Fail(fail types.FailProcessorVariant) *_processorContainer {

	s.v.Fail = fail.FailProcessorCaster()

	return s
}

func (s *_processorContainer) Fingerprint(fingerprint types.FingerprintProcessorVariant) *_processorContainer {

	s.v.Fingerprint = fingerprint.FingerprintProcessorCaster()

	return s
}

func (s *_processorContainer) Foreach(foreach types.ForeachProcessorVariant) *_processorContainer {

	s.v.Foreach = foreach.ForeachProcessorCaster()

	return s
}

func (s *_processorContainer) GeoGrid(geogrid types.GeoGridProcessorVariant) *_processorContainer {

	s.v.GeoGrid = geogrid.GeoGridProcessorCaster()

	return s
}

func (s *_processorContainer) Geoip(geoip types.GeoIpProcessorVariant) *_processorContainer {

	s.v.Geoip = geoip.GeoIpProcessorCaster()

	return s
}

func (s *_processorContainer) Grok(grok types.GrokProcessorVariant) *_processorContainer {

	s.v.Grok = grok.GrokProcessorCaster()

	return s
}

func (s *_processorContainer) Gsub(gsub types.GsubProcessorVariant) *_processorContainer {

	s.v.Gsub = gsub.GsubProcessorCaster()

	return s
}

func (s *_processorContainer) HtmlStrip(htmlstrip types.HtmlStripProcessorVariant) *_processorContainer {

	s.v.HtmlStrip = htmlstrip.HtmlStripProcessorCaster()

	return s
}

func (s *_processorContainer) Inference(inference types.InferenceProcessorVariant) *_processorContainer {

	s.v.Inference = inference.InferenceProcessorCaster()

	return s
}

func (s *_processorContainer) IpLocation(iplocation types.IpLocationProcessorVariant) *_processorContainer {

	s.v.IpLocation = iplocation.IpLocationProcessorCaster()

	return s
}

func (s *_processorContainer) Join(join types.JoinProcessorVariant) *_processorContainer {

	s.v.Join = join.JoinProcessorCaster()

	return s
}

func (s *_processorContainer) Json(json types.JsonProcessorVariant) *_processorContainer {

	s.v.Json = json.JsonProcessorCaster()

	return s
}

func (s *_processorContainer) Kv(kv types.KeyValueProcessorVariant) *_processorContainer {

	s.v.Kv = kv.KeyValueProcessorCaster()

	return s
}

func (s *_processorContainer) Lowercase(lowercase types.LowercaseProcessorVariant) *_processorContainer {

	s.v.Lowercase = lowercase.LowercaseProcessorCaster()

	return s
}

func (s *_processorContainer) NetworkDirection(networkdirection types.NetworkDirectionProcessorVariant) *_processorContainer {

	s.v.NetworkDirection = networkdirection.NetworkDirectionProcessorCaster()

	return s
}

func (s *_processorContainer) Pipeline(pipeline types.PipelineProcessorVariant) *_processorContainer {

	s.v.Pipeline = pipeline.PipelineProcessorCaster()

	return s
}

func (s *_processorContainer) Redact(redact types.RedactProcessorVariant) *_processorContainer {

	s.v.Redact = redact.RedactProcessorCaster()

	return s
}

func (s *_processorContainer) RegisteredDomain(registereddomain types.RegisteredDomainProcessorVariant) *_processorContainer {

	s.v.RegisteredDomain = registereddomain.RegisteredDomainProcessorCaster()

	return s
}

func (s *_processorContainer) Remove(remove types.RemoveProcessorVariant) *_processorContainer {

	s.v.Remove = remove.RemoveProcessorCaster()

	return s
}

func (s *_processorContainer) Rename(rename types.RenameProcessorVariant) *_processorContainer {

	s.v.Rename = rename.RenameProcessorCaster()

	return s
}

func (s *_processorContainer) Reroute(reroute types.RerouteProcessorVariant) *_processorContainer {

	s.v.Reroute = reroute.RerouteProcessorCaster()

	return s
}

func (s *_processorContainer) Script(script types.ScriptProcessorVariant) *_processorContainer {

	s.v.Script = script.ScriptProcessorCaster()

	return s
}

func (s *_processorContainer) Set(set types.SetProcessorVariant) *_processorContainer {

	s.v.Set = set.SetProcessorCaster()

	return s
}

func (s *_processorContainer) SetSecurityUser(setsecurityuser types.SetSecurityUserProcessorVariant) *_processorContainer {

	s.v.SetSecurityUser = setsecurityuser.SetSecurityUserProcessorCaster()

	return s
}

func (s *_processorContainer) Sort(sort types.SortProcessorVariant) *_processorContainer {

	s.v.Sort = sort.SortProcessorCaster()

	return s
}

func (s *_processorContainer) Split(split types.SplitProcessorVariant) *_processorContainer {

	s.v.Split = split.SplitProcessorCaster()

	return s
}

func (s *_processorContainer) Terminate(terminate types.TerminateProcessorVariant) *_processorContainer {

	s.v.Terminate = terminate.TerminateProcessorCaster()

	return s
}

func (s *_processorContainer) Trim(trim types.TrimProcessorVariant) *_processorContainer {

	s.v.Trim = trim.TrimProcessorCaster()

	return s
}

func (s *_processorContainer) Uppercase(uppercase types.UppercaseProcessorVariant) *_processorContainer {

	s.v.Uppercase = uppercase.UppercaseProcessorCaster()

	return s
}

func (s *_processorContainer) UriParts(uriparts types.UriPartsProcessorVariant) *_processorContainer {

	s.v.UriParts = uriparts.UriPartsProcessorCaster()

	return s
}

func (s *_processorContainer) Urldecode(urldecode types.UrlDecodeProcessorVariant) *_processorContainer {

	s.v.Urldecode = urldecode.UrlDecodeProcessorCaster()

	return s
}

func (s *_processorContainer) UserAgent(useragent types.UserAgentProcessorVariant) *_processorContainer {

	s.v.UserAgent = useragent.UserAgentProcessorCaster()

	return s
}

func (s *_processorContainer) ProcessorContainerCaster() *types.ProcessorContainer {
	return s.v
}
