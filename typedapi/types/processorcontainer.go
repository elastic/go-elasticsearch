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

package types

import (
	"encoding/json"
	"fmt"
)

// ProcessorContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L27-L301
type ProcessorContainer struct {
	AdditionalProcessorContainerProperty map[string]json.RawMessage `json:"-"`
	// Append Appends one or more values to an existing array if the field already exists
	// and it is an array.
	// Converts a scalar to an array and appends one or more values to it if the
	// field exists and it is a scalar.
	// Creates an array containing the provided values if the field doesn’t exist.
	// Accepts a single value or an array of values.
	Append *AppendProcessor `json:"append,omitempty"`
	// Attachment The attachment processor lets Elasticsearch extract file attachments in
	// common formats (such as PPT, XLS, and PDF) by using the Apache text
	// extraction library Tika.
	Attachment *AttachmentProcessor `json:"attachment,omitempty"`
	// Bytes Converts a human readable byte value (for example `1kb`) to its value in
	// bytes (for example `1024`).
	// If the field is an array of strings, all members of the array will be
	// converted.
	// Supported human readable units are "b", "kb", "mb", "gb", "tb", "pb" case
	// insensitive.
	// An error will occur if the field is not a supported format or resultant value
	// exceeds 2^63.
	Bytes *BytesProcessor `json:"bytes,omitempty"`
	// Circle Converts circle definitions of shapes to regular polygons which approximate
	// them.
	Circle *CircleProcessor `json:"circle,omitempty"`
	// CommunityId Computes the Community ID for network flow data as defined in the
	// Community ID Specification. You can use a community ID to correlate network
	// events related to a single flow.
	CommunityId *CommunityIDProcessor `json:"community_id,omitempty"`
	// Convert Converts a field in the currently ingested document to a different type, such
	// as converting a string to an integer.
	// If the field value is an array, all members will be converted.
	Convert *ConvertProcessor `json:"convert,omitempty"`
	// Csv Extracts fields from CSV line out of a single text field within a document.
	// Any empty field in CSV will be skipped.
	Csv *CsvProcessor `json:"csv,omitempty"`
	// Date Parses dates from fields, and then uses the date or timestamp as the
	// timestamp for the document.
	Date *DateProcessor `json:"date,omitempty"`
	// DateIndexName The purpose of this processor is to point documents to the right time based
	// index based on a date or timestamp field in a document by using the date math
	// index name support.
	DateIndexName *DateIndexNameProcessor `json:"date_index_name,omitempty"`
	// Dissect Extracts structured fields out of a single text field by matching the text
	// field against a delimiter-based pattern.
	Dissect *DissectProcessor `json:"dissect,omitempty"`
	// DotExpander Expands a field with dots into an object field.
	// This processor allows fields with dots in the name to be accessible by other
	// processors in the pipeline.
	// Otherwise these fields can’t be accessed by any processor.
	DotExpander *DotExpanderProcessor `json:"dot_expander,omitempty"`
	// Drop Drops the document without raising any errors.
	// This is useful to prevent the document from getting indexed based on some
	// condition.
	Drop *DropProcessor `json:"drop,omitempty"`
	// Enrich The `enrich` processor can enrich documents with data from another index.
	Enrich *EnrichProcessor `json:"enrich,omitempty"`
	// Fail Raises an exception.
	// This is useful for when you expect a pipeline to fail and want to relay a
	// specific message to the requester.
	Fail *FailProcessor `json:"fail,omitempty"`
	// Fingerprint Computes a hash of the document’s content. You can use this hash for
	// content fingerprinting.
	Fingerprint *FingerprintProcessor `json:"fingerprint,omitempty"`
	// Foreach Runs an ingest processor on each element of an array or object.
	Foreach *ForeachProcessor `json:"foreach,omitempty"`
	// GeoGrid Converts geo-grid definitions of grid tiles or cells to regular bounding
	// boxes or polygons which describe their shape.
	// This is useful if there is a need to interact with the tile shapes as
	// spatially indexable fields.
	GeoGrid *GeoGridProcessor `json:"geo_grid,omitempty"`
	// Geoip The `geoip` processor adds information about the geographical location of an
	// IPv4 or IPv6 address.
	Geoip *GeoIpProcessor `json:"geoip,omitempty"`
	// Grok Extracts structured fields out of a single text field within a document.
	// You choose which field to extract matched fields from, as well as the grok
	// pattern you expect will match.
	// A grok pattern is like a regular expression that supports aliased expressions
	// that can be reused.
	Grok *GrokProcessor `json:"grok,omitempty"`
	// Gsub Converts a string field by applying a regular expression and a replacement.
	// If the field is an array of string, all members of the array will be
	// converted.
	// If any non-string values are encountered, the processor will throw an
	// exception.
	Gsub *GsubProcessor `json:"gsub,omitempty"`
	// HtmlStrip Removes HTML tags from the field.
	// If the field is an array of strings, HTML tags will be removed from all
	// members of the array.
	HtmlStrip *HtmlStripProcessor `json:"html_strip,omitempty"`
	// Inference Uses a pre-trained data frame analytics model or a model deployed for natural
	// language processing tasks to infer against the data that is being ingested in
	// the pipeline.
	Inference *InferenceProcessor `json:"inference,omitempty"`
	// IpLocation Currently an undocumented alias for GeoIP Processor.
	IpLocation *IpLocationProcessor `json:"ip_location,omitempty"`
	// Join Joins each element of an array into a single string using a separator
	// character between each element.
	// Throws an error when the field is not an array.
	Join *JoinProcessor `json:"join,omitempty"`
	// Json Converts a JSON string into a structured JSON object.
	Json *JsonProcessor `json:"json,omitempty"`
	// Kv This processor helps automatically parse messages (or specific event fields)
	// which are of the `foo=bar` variety.
	Kv *KeyValueProcessor `json:"kv,omitempty"`
	// Lowercase Converts a string to its lowercase equivalent.
	// If the field is an array of strings, all members of the array will be
	// converted.
	Lowercase *LowercaseProcessor `json:"lowercase,omitempty"`
	// NetworkDirection Calculates the network direction given a source IP address, destination IP
	// address, and a list of internal networks.
	NetworkDirection *NetworkDirectionProcessor `json:"network_direction,omitempty"`
	// Pipeline Executes another pipeline.
	Pipeline *PipelineProcessor `json:"pipeline,omitempty"`
	// Redact The Redact processor uses the Grok rules engine to obscure text in the input
	// document matching the given Grok patterns.
	// The processor can be used to obscure Personal Identifying Information (PII)
	// by configuring it to detect known patterns such as email or IP addresses.
	// Text that matches a Grok pattern is replaced with a configurable string such
	// as `<EMAIL>` where an email address is matched or simply replace all matches
	// with the text `<REDACTED>` if preferred.
	Redact *RedactProcessor `json:"redact,omitempty"`
	// RegisteredDomain Extracts the registered domain (also known as the effective top-level
	// domain or eTLD), sub-domain, and top-level domain from a fully qualified
	// domain name (FQDN). Uses the registered domains defined in the Mozilla
	// Public Suffix List.
	RegisteredDomain *RegisteredDomainProcessor `json:"registered_domain,omitempty"`
	// Remove Removes existing fields.
	// If one field doesn’t exist, an exception will be thrown.
	Remove *RemoveProcessor `json:"remove,omitempty"`
	// Rename Renames an existing field.
	// If the field doesn’t exist or the new name is already used, an exception will
	// be thrown.
	Rename *RenameProcessor `json:"rename,omitempty"`
	// Reroute Routes a document to another target index or data stream.
	// When setting the `destination` option, the target is explicitly specified and
	// the dataset and namespace options can’t be set.
	// When the `destination` option is not set, this processor is in a data stream
	// mode. Note that in this mode, the reroute processor can only be used on data
	// streams that follow the data stream naming scheme.
	Reroute *RerouteProcessor `json:"reroute,omitempty"`
	// Script Runs an inline or stored script on incoming documents.
	// The script runs in the `ingest` context.
	Script *ScriptProcessor `json:"script,omitempty"`
	// Set Adds a field with the specified value.
	// If the field already exists, its value will be replaced with the provided
	// one.
	Set *SetProcessor `json:"set,omitempty"`
	// SetSecurityUser Sets user-related details (such as `username`, `roles`, `email`, `full_name`,
	// `metadata`, `api_key`, `realm` and `authentication_type`) from the current
	// authenticated user to the current document by pre-processing the ingest.
	SetSecurityUser *SetSecurityUserProcessor `json:"set_security_user,omitempty"`
	// Sort Sorts the elements of an array ascending or descending.
	// Homogeneous arrays of numbers will be sorted numerically, while arrays of
	// strings or heterogeneous arrays of strings + numbers will be sorted
	// lexicographically.
	// Throws an error when the field is not an array.
	Sort *SortProcessor `json:"sort,omitempty"`
	// Split Splits a field into an array using a separator character.
	// Only works on string fields.
	Split *SplitProcessor `json:"split,omitempty"`
	// Terminate Terminates the current ingest pipeline, causing no further processors to be
	// run.
	// This will normally be executed conditionally, using the `if` option.
	Terminate *TerminateProcessor `json:"terminate,omitempty"`
	// Trim Trims whitespace from a field.
	// If the field is an array of strings, all members of the array will be
	// trimmed.
	// This only works on leading and trailing whitespace.
	Trim *TrimProcessor `json:"trim,omitempty"`
	// Uppercase Converts a string to its uppercase equivalent.
	// If the field is an array of strings, all members of the array will be
	// converted.
	Uppercase *UppercaseProcessor `json:"uppercase,omitempty"`
	// UriParts Parses a Uniform Resource Identifier (URI) string and extracts its components
	// as an object.
	// This URI object includes properties for the URI’s domain, path, fragment,
	// port, query, scheme, user info, username, and password.
	UriParts *UriPartsProcessor `json:"uri_parts,omitempty"`
	// Urldecode URL-decodes a string.
	// If the field is an array of strings, all members of the array will be
	// decoded.
	Urldecode *UrlDecodeProcessor `json:"urldecode,omitempty"`
	// UserAgent The `user_agent` processor extracts details from the user agent string a
	// browser sends with its web requests.
	// This processor adds this information by default under the `user_agent` field.
	UserAgent *UserAgentProcessor `json:"user_agent,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ProcessorContainer) MarshalJSON() ([]byte, error) {
	type opt ProcessorContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalProcessorContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalProcessorContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewProcessorContainer returns a ProcessorContainer.
func NewProcessorContainer() *ProcessorContainer {
	r := &ProcessorContainer{
		AdditionalProcessorContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
