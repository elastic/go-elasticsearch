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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// ProcessorContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ingest/_types/Processors.ts#L28-L233
type ProcessorContainer struct {
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
	// Foreach Runs an ingest processor on each element of an array or object.
	Foreach *ForeachProcessor `json:"foreach,omitempty"`
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
	// Inference Uses a pre-trained data frame analytics model or a model deployed for natural
	// language processing tasks to infer against the data that is being ingested in
	// the pipeline.
	Inference *InferenceProcessor `json:"inference,omitempty"`
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
	// Pipeline Executes another pipeline.
	Pipeline *PipelineProcessor `json:"pipeline,omitempty"`
	// Remove Removes existing fields.
	// If one field doesn’t exist, an exception will be thrown.
	Remove *RemoveProcessor `json:"remove,omitempty"`
	// Rename Renames an existing field.
	// If the field doesn’t exist or the new name is already used, an exception will
	// be thrown.
	Rename *RenameProcessor `json:"rename,omitempty"`
	// Script Runs an inline or stored script on incoming documents.
	// The script runs in the `ingest` context.
	Script Script `json:"script,omitempty"`
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
	// Trim Trims whitespace from a field.
	// If the field is an array of strings, all members of the array will be
	// trimmed.
	// This only works on leading and trailing whitespace.
	Trim *TrimProcessor `json:"trim,omitempty"`
	// Uppercase Converts a string to its uppercase equivalent.
	// If the field is an array of strings, all members of the array will be
	// converted.
	Uppercase *UppercaseProcessor `json:"uppercase,omitempty"`
	// Urldecode URL-decodes a string.
	// If the field is an array of strings, all members of the array will be
	// decoded.
	Urldecode *UrlDecodeProcessor `json:"urldecode,omitempty"`
	// UserAgent The `user_agent` processor extracts details from the user agent string a
	// browser sends with its web requests.
	// This processor adds this information by default under the `user_agent` field.
	UserAgent *UserAgentProcessor `json:"user_agent,omitempty"`
}

func (s *ProcessorContainer) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "append":
			if err := dec.Decode(&s.Append); err != nil {
				return err
			}

		case "attachment":
			if err := dec.Decode(&s.Attachment); err != nil {
				return err
			}

		case "bytes":
			if err := dec.Decode(&s.Bytes); err != nil {
				return err
			}

		case "circle":
			if err := dec.Decode(&s.Circle); err != nil {
				return err
			}

		case "convert":
			if err := dec.Decode(&s.Convert); err != nil {
				return err
			}

		case "csv":
			if err := dec.Decode(&s.Csv); err != nil {
				return err
			}

		case "date":
			if err := dec.Decode(&s.Date); err != nil {
				return err
			}

		case "date_index_name":
			if err := dec.Decode(&s.DateIndexName); err != nil {
				return err
			}

		case "dissect":
			if err := dec.Decode(&s.Dissect); err != nil {
				return err
			}

		case "dot_expander":
			if err := dec.Decode(&s.DotExpander); err != nil {
				return err
			}

		case "drop":
			if err := dec.Decode(&s.Drop); err != nil {
				return err
			}

		case "enrich":
			if err := dec.Decode(&s.Enrich); err != nil {
				return err
			}

		case "fail":
			if err := dec.Decode(&s.Fail); err != nil {
				return err
			}

		case "foreach":
			if err := dec.Decode(&s.Foreach); err != nil {
				return err
			}

		case "geoip":
			if err := dec.Decode(&s.Geoip); err != nil {
				return err
			}

		case "grok":
			if err := dec.Decode(&s.Grok); err != nil {
				return err
			}

		case "gsub":
			if err := dec.Decode(&s.Gsub); err != nil {
				return err
			}

		case "inference":
			if err := dec.Decode(&s.Inference); err != nil {
				return err
			}

		case "join":
			if err := dec.Decode(&s.Join); err != nil {
				return err
			}

		case "json":
			if err := dec.Decode(&s.Json); err != nil {
				return err
			}

		case "kv":
			if err := dec.Decode(&s.Kv); err != nil {
				return err
			}

		case "lowercase":
			if err := dec.Decode(&s.Lowercase); err != nil {
				return err
			}

		case "pipeline":
			if err := dec.Decode(&s.Pipeline); err != nil {
				return err
			}

		case "remove":
			if err := dec.Decode(&s.Remove); err != nil {
				return err
			}

		case "rename":
			if err := dec.Decode(&s.Rename); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "set":
			if err := dec.Decode(&s.Set); err != nil {
				return err
			}

		case "set_security_user":
			if err := dec.Decode(&s.SetSecurityUser); err != nil {
				return err
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return err
			}

		case "split":
			if err := dec.Decode(&s.Split); err != nil {
				return err
			}

		case "trim":
			if err := dec.Decode(&s.Trim); err != nil {
				return err
			}

		case "uppercase":
			if err := dec.Decode(&s.Uppercase); err != nil {
				return err
			}

		case "urldecode":
			if err := dec.Decode(&s.Urldecode); err != nil {
				return err
			}

		case "user_agent":
			if err := dec.Decode(&s.UserAgent); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewProcessorContainer returns a ProcessorContainer.
func NewProcessorContainer() *ProcessorContainer {
	r := &ProcessorContainer{}

	return r
}
