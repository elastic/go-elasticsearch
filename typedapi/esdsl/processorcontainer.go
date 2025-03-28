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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

// Appends one or more values to an existing array if the field already exists
// and it is an array.
// Converts a scalar to an array and appends one or more values to it if the
// field exists and it is a scalar.
// Creates an array containing the provided values if the field doesn’t exist.
// Accepts a single value or an array of values.
func (s *_processorContainer) Append(append types.AppendProcessorVariant) *_processorContainer {

	s.v.Append = append.AppendProcessorCaster()

	return s
}

// The attachment processor lets Elasticsearch extract file attachments in
// common formats (such as PPT, XLS, and PDF) by using the Apache text
// extraction library Tika.
func (s *_processorContainer) Attachment(attachment types.AttachmentProcessorVariant) *_processorContainer {

	s.v.Attachment = attachment.AttachmentProcessorCaster()

	return s
}

// Converts a human readable byte value (for example `1kb`) to its value in
// bytes (for example `1024`).
// If the field is an array of strings, all members of the array will be
// converted.
// Supported human readable units are "b", "kb", "mb", "gb", "tb", "pb" case
// insensitive.
// An error will occur if the field is not a supported format or resultant value
// exceeds 2^63.
func (s *_processorContainer) Bytes(bytes types.BytesProcessorVariant) *_processorContainer {

	s.v.Bytes = bytes.BytesProcessorCaster()

	return s
}

// Converts circle definitions of shapes to regular polygons which approximate
// them.
func (s *_processorContainer) Circle(circle types.CircleProcessorVariant) *_processorContainer {

	s.v.Circle = circle.CircleProcessorCaster()

	return s
}

// Computes the Community ID for network flow data as defined in the
// Community ID Specification. You can use a community ID to correlate network
// events related to a single flow.
func (s *_processorContainer) CommunityId(communityid types.CommunityIDProcessorVariant) *_processorContainer {

	s.v.CommunityId = communityid.CommunityIDProcessorCaster()

	return s
}

// Converts a field in the currently ingested document to a different type, such
// as converting a string to an integer.
// If the field value is an array, all members will be converted.
func (s *_processorContainer) Convert(convert types.ConvertProcessorVariant) *_processorContainer {

	s.v.Convert = convert.ConvertProcessorCaster()

	return s
}

// Extracts fields from CSV line out of a single text field within a document.
// Any empty field in CSV will be skipped.
func (s *_processorContainer) Csv(csv types.CsvProcessorVariant) *_processorContainer {

	s.v.Csv = csv.CsvProcessorCaster()

	return s
}

// Parses dates from fields, and then uses the date or timestamp as the
// timestamp for the document.
func (s *_processorContainer) Date(date types.DateProcessorVariant) *_processorContainer {

	s.v.Date = date.DateProcessorCaster()

	return s
}

// The purpose of this processor is to point documents to the right time based
// index based on a date or timestamp field in a document by using the date math
// index name support.
func (s *_processorContainer) DateIndexName(dateindexname types.DateIndexNameProcessorVariant) *_processorContainer {

	s.v.DateIndexName = dateindexname.DateIndexNameProcessorCaster()

	return s
}

// Extracts structured fields out of a single text field by matching the text
// field against a delimiter-based pattern.
func (s *_processorContainer) Dissect(dissect types.DissectProcessorVariant) *_processorContainer {

	s.v.Dissect = dissect.DissectProcessorCaster()

	return s
}

// Expands a field with dots into an object field.
// This processor allows fields with dots in the name to be accessible by other
// processors in the pipeline.
// Otherwise these fields can’t be accessed by any processor.
func (s *_processorContainer) DotExpander(dotexpander types.DotExpanderProcessorVariant) *_processorContainer {

	s.v.DotExpander = dotexpander.DotExpanderProcessorCaster()

	return s
}

// Drops the document without raising any errors.
// This is useful to prevent the document from getting indexed based on some
// condition.
func (s *_processorContainer) Drop(drop types.DropProcessorVariant) *_processorContainer {

	s.v.Drop = drop.DropProcessorCaster()

	return s
}

// The `enrich` processor can enrich documents with data from another index.
func (s *_processorContainer) Enrich(enrich types.EnrichProcessorVariant) *_processorContainer {

	s.v.Enrich = enrich.EnrichProcessorCaster()

	return s
}

// Raises an exception.
// This is useful for when you expect a pipeline to fail and want to relay a
// specific message to the requester.
func (s *_processorContainer) Fail(fail types.FailProcessorVariant) *_processorContainer {

	s.v.Fail = fail.FailProcessorCaster()

	return s
}

// Computes a hash of the document’s content. You can use this hash for
// content fingerprinting.
func (s *_processorContainer) Fingerprint(fingerprint types.FingerprintProcessorVariant) *_processorContainer {

	s.v.Fingerprint = fingerprint.FingerprintProcessorCaster()

	return s
}

// Runs an ingest processor on each element of an array or object.
func (s *_processorContainer) Foreach(foreach types.ForeachProcessorVariant) *_processorContainer {

	s.v.Foreach = foreach.ForeachProcessorCaster()

	return s
}

// Converts geo-grid definitions of grid tiles or cells to regular bounding
// boxes or polygons which describe their shape.
// This is useful if there is a need to interact with the tile shapes as
// spatially indexable fields.
func (s *_processorContainer) GeoGrid(geogrid types.GeoGridProcessorVariant) *_processorContainer {

	s.v.GeoGrid = geogrid.GeoGridProcessorCaster()

	return s
}

// The `geoip` processor adds information about the geographical location of an
// IPv4 or IPv6 address.
func (s *_processorContainer) Geoip(geoip types.GeoIpProcessorVariant) *_processorContainer {

	s.v.Geoip = geoip.GeoIpProcessorCaster()

	return s
}

// Extracts structured fields out of a single text field within a document.
// You choose which field to extract matched fields from, as well as the grok
// pattern you expect will match.
// A grok pattern is like a regular expression that supports aliased expressions
// that can be reused.
func (s *_processorContainer) Grok(grok types.GrokProcessorVariant) *_processorContainer {

	s.v.Grok = grok.GrokProcessorCaster()

	return s
}

// Converts a string field by applying a regular expression and a replacement.
// If the field is an array of string, all members of the array will be
// converted.
// If any non-string values are encountered, the processor will throw an
// exception.
func (s *_processorContainer) Gsub(gsub types.GsubProcessorVariant) *_processorContainer {

	s.v.Gsub = gsub.GsubProcessorCaster()

	return s
}

// Removes HTML tags from the field.
// If the field is an array of strings, HTML tags will be removed from all
// members of the array.
func (s *_processorContainer) HtmlStrip(htmlstrip types.HtmlStripProcessorVariant) *_processorContainer {

	s.v.HtmlStrip = htmlstrip.HtmlStripProcessorCaster()

	return s
}

// Uses a pre-trained data frame analytics model or a model deployed for natural
// language processing tasks to infer against the data that is being ingested in
// the pipeline.
func (s *_processorContainer) Inference(inference types.InferenceProcessorVariant) *_processorContainer {

	s.v.Inference = inference.InferenceProcessorCaster()

	return s
}

// Currently an undocumented alias for GeoIP Processor.
func (s *_processorContainer) IpLocation(iplocation types.IpLocationProcessorVariant) *_processorContainer {

	s.v.IpLocation = iplocation.IpLocationProcessorCaster()

	return s
}

// Joins each element of an array into a single string using a separator
// character between each element.
// Throws an error when the field is not an array.
func (s *_processorContainer) Join(join types.JoinProcessorVariant) *_processorContainer {

	s.v.Join = join.JoinProcessorCaster()

	return s
}

// Converts a JSON string into a structured JSON object.
func (s *_processorContainer) Json(json types.JsonProcessorVariant) *_processorContainer {

	s.v.Json = json.JsonProcessorCaster()

	return s
}

// This processor helps automatically parse messages (or specific event fields)
// which are of the `foo=bar` variety.
func (s *_processorContainer) Kv(kv types.KeyValueProcessorVariant) *_processorContainer {

	s.v.Kv = kv.KeyValueProcessorCaster()

	return s
}

// Converts a string to its lowercase equivalent.
// If the field is an array of strings, all members of the array will be
// converted.
func (s *_processorContainer) Lowercase(lowercase types.LowercaseProcessorVariant) *_processorContainer {

	s.v.Lowercase = lowercase.LowercaseProcessorCaster()

	return s
}

// Calculates the network direction given a source IP address, destination IP
// address, and a list of internal networks.
func (s *_processorContainer) NetworkDirection(networkdirection types.NetworkDirectionProcessorVariant) *_processorContainer {

	s.v.NetworkDirection = networkdirection.NetworkDirectionProcessorCaster()

	return s
}

// Executes another pipeline.
func (s *_processorContainer) Pipeline(pipeline types.PipelineProcessorVariant) *_processorContainer {

	s.v.Pipeline = pipeline.PipelineProcessorCaster()

	return s
}

// The Redact processor uses the Grok rules engine to obscure text in the input
// document matching the given Grok patterns.
// The processor can be used to obscure Personal Identifying Information (PII)
// by configuring it to detect known patterns such as email or IP addresses.
// Text that matches a Grok pattern is replaced with a configurable string such
// as `<EMAIL>` where an email address is matched or simply replace all matches
// with the text `<REDACTED>` if preferred.
func (s *_processorContainer) Redact(redact types.RedactProcessorVariant) *_processorContainer {

	s.v.Redact = redact.RedactProcessorCaster()

	return s
}

// Extracts the registered domain (also known as the effective top-level
// domain or eTLD), sub-domain, and top-level domain from a fully qualified
// domain name (FQDN). Uses the registered domains defined in the Mozilla
// Public Suffix List.
func (s *_processorContainer) RegisteredDomain(registereddomain types.RegisteredDomainProcessorVariant) *_processorContainer {

	s.v.RegisteredDomain = registereddomain.RegisteredDomainProcessorCaster()

	return s
}

// Removes existing fields.
// If one field doesn’t exist, an exception will be thrown.
func (s *_processorContainer) Remove(remove types.RemoveProcessorVariant) *_processorContainer {

	s.v.Remove = remove.RemoveProcessorCaster()

	return s
}

// Renames an existing field.
// If the field doesn’t exist or the new name is already used, an exception will
// be thrown.
func (s *_processorContainer) Rename(rename types.RenameProcessorVariant) *_processorContainer {

	s.v.Rename = rename.RenameProcessorCaster()

	return s
}

// Routes a document to another target index or data stream.
// When setting the `destination` option, the target is explicitly specified and
// the dataset and namespace options can’t be set.
// When the `destination` option is not set, this processor is in a data stream
// mode. Note that in this mode, the reroute processor can only be used on data
// streams that follow the data stream naming scheme.
func (s *_processorContainer) Reroute(reroute types.RerouteProcessorVariant) *_processorContainer {

	s.v.Reroute = reroute.RerouteProcessorCaster()

	return s
}

// Runs an inline or stored script on incoming documents.
// The script runs in the `ingest` context.
func (s *_processorContainer) Script(script types.ScriptProcessorVariant) *_processorContainer {

	s.v.Script = script.ScriptProcessorCaster()

	return s
}

// Adds a field with the specified value.
// If the field already exists, its value will be replaced with the provided
// one.
func (s *_processorContainer) Set(set types.SetProcessorVariant) *_processorContainer {

	s.v.Set = set.SetProcessorCaster()

	return s
}

// Sets user-related details (such as `username`, `roles`, `email`, `full_name`,
// `metadata`, `api_key`, `realm` and `authentication_type`) from the current
// authenticated user to the current document by pre-processing the ingest.
func (s *_processorContainer) SetSecurityUser(setsecurityuser types.SetSecurityUserProcessorVariant) *_processorContainer {

	s.v.SetSecurityUser = setsecurityuser.SetSecurityUserProcessorCaster()

	return s
}

// Sorts the elements of an array ascending or descending.
// Homogeneous arrays of numbers will be sorted numerically, while arrays of
// strings or heterogeneous arrays of strings + numbers will be sorted
// lexicographically.
// Throws an error when the field is not an array.
func (s *_processorContainer) Sort(sort types.SortProcessorVariant) *_processorContainer {

	s.v.Sort = sort.SortProcessorCaster()

	return s
}

// Splits a field into an array using a separator character.
// Only works on string fields.
func (s *_processorContainer) Split(split types.SplitProcessorVariant) *_processorContainer {

	s.v.Split = split.SplitProcessorCaster()

	return s
}

// Terminates the current ingest pipeline, causing no further processors to be
// run.
// This will normally be executed conditionally, using the `if` option.
func (s *_processorContainer) Terminate(terminate types.TerminateProcessorVariant) *_processorContainer {

	s.v.Terminate = terminate.TerminateProcessorCaster()

	return s
}

// Trims whitespace from a field.
// If the field is an array of strings, all members of the array will be
// trimmed.
// This only works on leading and trailing whitespace.
func (s *_processorContainer) Trim(trim types.TrimProcessorVariant) *_processorContainer {

	s.v.Trim = trim.TrimProcessorCaster()

	return s
}

// Converts a string to its uppercase equivalent.
// If the field is an array of strings, all members of the array will be
// converted.
func (s *_processorContainer) Uppercase(uppercase types.UppercaseProcessorVariant) *_processorContainer {

	s.v.Uppercase = uppercase.UppercaseProcessorCaster()

	return s
}

// Parses a Uniform Resource Identifier (URI) string and extracts its components
// as an object.
// This URI object includes properties for the URI’s domain, path, fragment,
// port, query, scheme, user info, username, and password.
func (s *_processorContainer) UriParts(uriparts types.UriPartsProcessorVariant) *_processorContainer {

	s.v.UriParts = uriparts.UriPartsProcessorCaster()

	return s
}

// URL-decodes a string.
// If the field is an array of strings, all members of the array will be
// decoded.
func (s *_processorContainer) Urldecode(urldecode types.UrlDecodeProcessorVariant) *_processorContainer {

	s.v.Urldecode = urldecode.UrlDecodeProcessorCaster()

	return s
}

// The `user_agent` processor extracts details from the user agent string a
// browser sends with its web requests.
// This processor adds this information by default under the `user_agent` field.
func (s *_processorContainer) UserAgent(useragent types.UserAgentProcessorVariant) *_processorContainer {

	s.v.UserAgent = useragent.UserAgentProcessorCaster()

	return s
}

func (s *_processorContainer) ProcessorContainerCaster() *types.ProcessorContainer {
	return s.v
}
