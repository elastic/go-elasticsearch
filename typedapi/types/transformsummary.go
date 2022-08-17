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

// TransformSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/get_transform/types.ts#L33-L61
type TransformSummary struct {
	// Authorization The security privileges that the transform uses to run its queries. If
	// Elastic Stack security features were disabled at the time of the most recent
	// update to the transform, this property is omitted.
	Authorization *TransformAuthorization `json:"authorization,omitempty"`
	// CreateTime The time the transform was created.
	CreateTime *EpochTimeUnitMillis `json:"create_time,omitempty"`
	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest      Destination `json:"dest"`
	Frequency *Duration   `json:"frequency,omitempty"`
	Id        Id          `json:"id"`
	Latest    *Latest     `json:"latest,omitempty"`
	Meta_     *Metadata   `json:"_meta,omitempty"`
	// Pivot The pivot method transforms the data by aggregating and grouping it.
	Pivot           *Pivot                    `json:"pivot,omitempty"`
	RetentionPolicy *RetentionPolicyContainer `json:"retention_policy,omitempty"`
	// Settings Defines optional transform settings.
	Settings *Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source Source `json:"source"`
	// Sync Defines the properties transforms require to run continuously.
	Sync *SyncContainer `json:"sync,omitempty"`
	// Version The version of Elasticsearch that existed on the node when the transform was
	// created.
	Version *VersionString `json:"version,omitempty"`
}

// TransformSummaryBuilder holds TransformSummary struct and provides a builder API.
type TransformSummaryBuilder struct {
	v *TransformSummary
}

// NewTransformSummary provides a builder for the TransformSummary struct.
func NewTransformSummaryBuilder() *TransformSummaryBuilder {
	r := TransformSummaryBuilder{
		&TransformSummary{},
	}

	return &r
}

// Build finalize the chain and returns the TransformSummary struct
func (rb *TransformSummaryBuilder) Build() TransformSummary {
	return *rb.v
}

// Authorization The security privileges that the transform uses to run its queries. If
// Elastic Stack security features were disabled at the time of the most recent
// update to the transform, this property is omitted.

func (rb *TransformSummaryBuilder) Authorization(authorization *TransformAuthorizationBuilder) *TransformSummaryBuilder {
	v := authorization.Build()
	rb.v.Authorization = &v
	return rb
}

// CreateTime The time the transform was created.

func (rb *TransformSummaryBuilder) CreateTime(createtime *EpochTimeUnitMillisBuilder) *TransformSummaryBuilder {
	v := createtime.Build()
	rb.v.CreateTime = &v
	return rb
}

// Description Free text description of the transform.

func (rb *TransformSummaryBuilder) Description(description string) *TransformSummaryBuilder {
	rb.v.Description = &description
	return rb
}

// Dest The destination for the transform.

func (rb *TransformSummaryBuilder) Dest(dest *DestinationBuilder) *TransformSummaryBuilder {
	v := dest.Build()
	rb.v.Dest = v
	return rb
}

func (rb *TransformSummaryBuilder) Frequency(frequency *DurationBuilder) *TransformSummaryBuilder {
	v := frequency.Build()
	rb.v.Frequency = &v
	return rb
}

func (rb *TransformSummaryBuilder) Id(id Id) *TransformSummaryBuilder {
	rb.v.Id = id
	return rb
}

func (rb *TransformSummaryBuilder) Latest(latest *LatestBuilder) *TransformSummaryBuilder {
	v := latest.Build()
	rb.v.Latest = &v
	return rb
}

func (rb *TransformSummaryBuilder) Meta_(meta_ *MetadataBuilder) *TransformSummaryBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

// Pivot The pivot method transforms the data by aggregating and grouping it.

func (rb *TransformSummaryBuilder) Pivot(pivot *PivotBuilder) *TransformSummaryBuilder {
	v := pivot.Build()
	rb.v.Pivot = &v
	return rb
}

func (rb *TransformSummaryBuilder) RetentionPolicy(retentionpolicy *RetentionPolicyContainerBuilder) *TransformSummaryBuilder {
	v := retentionpolicy.Build()
	rb.v.RetentionPolicy = &v
	return rb
}

// Settings Defines optional transform settings.

func (rb *TransformSummaryBuilder) Settings(settings *SettingsBuilder) *TransformSummaryBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}

// Source The source of the data for the transform.

func (rb *TransformSummaryBuilder) Source(source *SourceBuilder) *TransformSummaryBuilder {
	v := source.Build()
	rb.v.Source = v
	return rb
}

// Sync Defines the properties transforms require to run continuously.

func (rb *TransformSummaryBuilder) Sync(sync *SyncContainerBuilder) *TransformSummaryBuilder {
	v := sync.Build()
	rb.v.Sync = &v
	return rb
}

// Version The version of Elasticsearch that existed on the node when the transform was
// created.

func (rb *TransformSummaryBuilder) Version(version VersionString) *TransformSummaryBuilder {
	rb.v.Version = &version
	return rb
}
