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

// Verify the repository integrity.
// Verify the integrity of the contents of a snapshot repository.
//
// This API enables you to perform a comprehensive check of the contents of a
// repository, looking for any anomalies in its data or metadata which might
// prevent you from restoring snapshots from the repository or which might cause
// future snapshot create or delete operations to fail.
//
// If you suspect the integrity of the contents of one of your snapshot
// repositories, cease all write activity to this repository immediately, set
// its `read_only` option to `true`, and use this API to verify its integrity.
// Until you do so:
//
// * It may not be possible to restore some snapshots from this repository.
// * Searchable snapshots may report errors when searched or may have unassigned
// shards.
// * Taking snapshots into this repository may fail or may appear to succeed but
// have created a snapshot which cannot be restored.
// * Deleting snapshots from this repository may fail or may appear to succeed
// but leave the underlying data on disk.
// * Continuing to write to the repository while it is in an invalid state may
// causing additional damage to its contents.
//
// If the API finds any problems with the integrity of the contents of your
// repository, Elasticsearch will not be able to repair the damage.
// The only way to bring the repository back into a fully working state after
// its contents have been damaged is by restoring its contents from a repository
// backup which was taken before the damage occurred.
// You must also identify what caused the damage and take action to prevent it
// from happening again.
//
// If you cannot restore a repository backup, register a new repository and use
// this for all future snapshot operations.
// In some cases it may be possible to recover some of the contents of a damaged
// repository, either by restoring as many of its snapshots as needed and taking
// new snapshots of the restored data, or by using the reindex API to copy data
// from any searchable snapshots mounted from the damaged repository.
//
// Avoid all operations which write to the repository while the verify
// repository integrity API is running.
// If something changes the repository contents while an integrity verification
// is running then Elasticsearch may incorrectly report having detected some
// anomalies in its contents due to the concurrent writes.
// It may also incorrectly fail to report some anomalies that the concurrent
// writes prevented it from detecting.
//
// NOTE: This API is intended for exploratory use by humans. You should expect
// the request parameters and the response format to vary in future versions.
//
// NOTE: This API may not work correctly in a mixed-version cluster.
package repositoryverifyintegrity

import (
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
)

const (
	repositoryMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RepositoryVerifyIntegrity struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewRepositoryVerifyIntegrity type alias for index.
type NewRepositoryVerifyIntegrity func(repository string) *RepositoryVerifyIntegrity

// NewRepositoryVerifyIntegrityFunc returns a new instance of RepositoryVerifyIntegrity with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRepositoryVerifyIntegrityFunc(tp elastictransport.Interface) NewRepositoryVerifyIntegrity {
	return func(repository string) *RepositoryVerifyIntegrity {
		n := New(tp)

		n._repository(repository)

		return n
	}
}

// Verify the repository integrity.
// Verify the integrity of the contents of a snapshot repository.
//
// This API enables you to perform a comprehensive check of the contents of a
// repository, looking for any anomalies in its data or metadata which might
// prevent you from restoring snapshots from the repository or which might cause
// future snapshot create or delete operations to fail.
//
// If you suspect the integrity of the contents of one of your snapshot
// repositories, cease all write activity to this repository immediately, set
// its `read_only` option to `true`, and use this API to verify its integrity.
// Until you do so:
//
// * It may not be possible to restore some snapshots from this repository.
// * Searchable snapshots may report errors when searched or may have unassigned
// shards.
// * Taking snapshots into this repository may fail or may appear to succeed but
// have created a snapshot which cannot be restored.
// * Deleting snapshots from this repository may fail or may appear to succeed
// but leave the underlying data on disk.
// * Continuing to write to the repository while it is in an invalid state may
// causing additional damage to its contents.
//
// If the API finds any problems with the integrity of the contents of your
// repository, Elasticsearch will not be able to repair the damage.
// The only way to bring the repository back into a fully working state after
// its contents have been damaged is by restoring its contents from a repository
// backup which was taken before the damage occurred.
// You must also identify what caused the damage and take action to prevent it
// from happening again.
//
// If you cannot restore a repository backup, register a new repository and use
// this for all future snapshot operations.
// In some cases it may be possible to recover some of the contents of a damaged
// repository, either by restoring as many of its snapshots as needed and taking
// new snapshots of the restored data, or by using the reindex API to copy data
// from any searchable snapshots mounted from the damaged repository.
//
// Avoid all operations which write to the repository while the verify
// repository integrity API is running.
// If something changes the repository contents while an integrity verification
// is running then Elasticsearch may incorrectly report having detected some
// anomalies in its contents due to the concurrent writes.
// It may also incorrectly fail to report some anomalies that the concurrent
// writes prevented it from detecting.
//
// NOTE: This API is intended for exploratory use by humans. You should expect
// the request parameters and the response format to vary in future versions.
//
// NOTE: This API may not work correctly in a mixed-version cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/verify-repo-integrity-api.html
func New(tp elastictransport.Interface) *RepositoryVerifyIntegrity {
	r := &RepositoryVerifyIntegrity{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *RepositoryVerifyIntegrity) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "repository", r.repository)
		}
		path.WriteString(r.repository)
		path.WriteString("/")
		path.WriteString("_verify_integrity")

		method = http.MethodPost
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r RepositoryVerifyIntegrity) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.repository_verify_integrity")
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
		instrument.BeforeRequest(req, "snapshot.repository_verify_integrity")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.repository_verify_integrity", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.repository_verify_integrity")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the RepositoryVerifyIntegrity query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a repositoryverifyintegrity.Response
func (r RepositoryVerifyIntegrity) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.repository_verify_integrity")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := new(Response)

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return *response, nil
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r RepositoryVerifyIntegrity) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.repository_verify_integrity")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the RepositoryVerifyIntegrity query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the RepositoryVerifyIntegrity headers map.
func (r *RepositoryVerifyIntegrity) Header(key, value string) *RepositoryVerifyIntegrity {
	r.headers.Set(key, value)

	return r
}

// Repository A repository name
// API Name: repository
func (r *RepositoryVerifyIntegrity) _repository(repository string) *RepositoryVerifyIntegrity {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// MetaThreadPoolConcurrency Number of threads to use for reading metadata
// API name: meta_thread_pool_concurrency
func (r *RepositoryVerifyIntegrity) MetaThreadPoolConcurrency(metathreadpoolconcurrency int) *RepositoryVerifyIntegrity {
	r.values.Set("meta_thread_pool_concurrency", strconv.Itoa(metathreadpoolconcurrency))

	return r
}

// BlobThreadPoolConcurrency Number of threads to use for reading blob contents
// API name: blob_thread_pool_concurrency
func (r *RepositoryVerifyIntegrity) BlobThreadPoolConcurrency(blobthreadpoolconcurrency int) *RepositoryVerifyIntegrity {
	r.values.Set("blob_thread_pool_concurrency", strconv.Itoa(blobthreadpoolconcurrency))

	return r
}

// SnapshotVerificationConcurrency Number of snapshots to verify concurrently
// API name: snapshot_verification_concurrency
func (r *RepositoryVerifyIntegrity) SnapshotVerificationConcurrency(snapshotverificationconcurrency int) *RepositoryVerifyIntegrity {
	r.values.Set("snapshot_verification_concurrency", strconv.Itoa(snapshotverificationconcurrency))

	return r
}

// IndexVerificationConcurrency Number of indices to verify concurrently
// API name: index_verification_concurrency
func (r *RepositoryVerifyIntegrity) IndexVerificationConcurrency(indexverificationconcurrency int) *RepositoryVerifyIntegrity {
	r.values.Set("index_verification_concurrency", strconv.Itoa(indexverificationconcurrency))

	return r
}

// IndexSnapshotVerificationConcurrency Number of snapshots to verify concurrently within each index
// API name: index_snapshot_verification_concurrency
func (r *RepositoryVerifyIntegrity) IndexSnapshotVerificationConcurrency(indexsnapshotverificationconcurrency int) *RepositoryVerifyIntegrity {
	r.values.Set("index_snapshot_verification_concurrency", strconv.Itoa(indexsnapshotverificationconcurrency))

	return r
}

// MaxFailedShardSnapshots Maximum permitted number of failed shard snapshots
// API name: max_failed_shard_snapshots
func (r *RepositoryVerifyIntegrity) MaxFailedShardSnapshots(maxfailedshardsnapshots int) *RepositoryVerifyIntegrity {
	r.values.Set("max_failed_shard_snapshots", strconv.Itoa(maxfailedshardsnapshots))

	return r
}

// VerifyBlobContents Whether to verify the contents of individual blobs
// API name: verify_blob_contents
func (r *RepositoryVerifyIntegrity) VerifyBlobContents(verifyblobcontents bool) *RepositoryVerifyIntegrity {
	r.values.Set("verify_blob_contents", strconv.FormatBool(verifyblobcontents))

	return r
}

// MaxBytesPerSec Rate limit for individual blob verification
// API name: max_bytes_per_sec
func (r *RepositoryVerifyIntegrity) MaxBytesPerSec(maxbytespersec string) *RepositoryVerifyIntegrity {
	r.values.Set("max_bytes_per_sec", maxbytespersec)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *RepositoryVerifyIntegrity) ErrorTrace(errortrace bool) *RepositoryVerifyIntegrity {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *RepositoryVerifyIntegrity) FilterPath(filterpaths ...string) *RepositoryVerifyIntegrity {
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
func (r *RepositoryVerifyIntegrity) Human(human bool) *RepositoryVerifyIntegrity {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *RepositoryVerifyIntegrity) Pretty(pretty bool) *RepositoryVerifyIntegrity {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
