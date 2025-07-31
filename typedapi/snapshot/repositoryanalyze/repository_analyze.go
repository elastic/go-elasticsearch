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

// Analyze a snapshot repository.
// Analyze the performance characteristics and any incorrect behaviour found in
// a repository.
//
// The response exposes implementation details of the analysis which may change
// from version to version.
// The response body format is therefore not considered stable and may be
// different in newer versions.
//
// There are a large number of third-party storage systems available, not all of
// which are suitable for use as a snapshot repository by Elasticsearch.
// Some storage systems behave incorrectly, or perform poorly, especially when
// accessed concurrently by multiple clients as the nodes of an Elasticsearch
// cluster do. This API performs a collection of read and write operations on
// your repository which are designed to detect incorrect behaviour and to
// measure the performance characteristics of your storage system.
//
// The default values for the parameters are deliberately low to reduce the
// impact of running an analysis inadvertently and to provide a sensible
// starting point for your investigations.
// Run your first analysis with the default parameter values to check for simple
// problems.
// If successful, run a sequence of increasingly large analyses until you
// encounter a failure or you reach a `blob_count` of at least `2000`, a
// `max_blob_size` of at least `2gb`, a `max_total_data_size` of at least `1tb`,
// and a `register_operation_count` of at least `100`.
// Always specify a generous timeout, possibly `1h` or longer, to allow time for
// each analysis to run to completion.
// Perform the analyses using a multi-node cluster of a similar size to your
// production cluster so that it can detect any problems that only arise when
// the repository is accessed by many nodes at once.
//
// If the analysis fails, Elasticsearch detected that your repository behaved
// unexpectedly.
// This usually means you are using a third-party storage system with an
// incorrect or incompatible implementation of the API it claims to support.
// If so, this storage system is not suitable for use as a snapshot repository.
// You will need to work with the supplier of your storage system to address the
// incompatibilities that Elasticsearch detects.
//
// If the analysis is successful, the API returns details of the testing
// process, optionally including how long each operation took.
// You can use this information to determine the performance of your storage
// system.
// If any operation fails or returns an incorrect result, the API returns an
// error.
// If the API returns an error, it may not have removed all the data it wrote to
// the repository.
// The error will indicate the location of any leftover data and this path is
// also recorded in the Elasticsearch logs.
// You should verify that this location has been cleaned up correctly.
// If there is still leftover data at the specified location, you should
// manually remove it.
//
// If the connection from your client to Elasticsearch is closed while the
// client is waiting for the result of the analysis, the test is cancelled.
// Some clients are configured to close their connection if no response is
// received within a certain timeout.
// An analysis takes a long time to complete so you might need to relax any such
// client-side timeouts.
// On cancellation the analysis attempts to clean up the data it was writing,
// but it may not be able to remove it all.
// The path to the leftover data is recorded in the Elasticsearch logs.
// You should verify that this location has been cleaned up correctly.
// If there is still leftover data at the specified location, you should
// manually remove it.
//
// If the analysis is successful then it detected no incorrect behaviour, but
// this does not mean that correct behaviour is guaranteed.
// The analysis attempts to detect common bugs but it does not offer 100%
// coverage.
// Additionally, it does not test the following:
//
// * Your repository must perform durable writes. Once a blob has been written
// it must remain in place until it is deleted, even after a power loss or
// similar disaster.
// * Your repository must not suffer from silent data corruption. Once a blob
// has been written, its contents must remain unchanged until it is deliberately
// modified or deleted.
// * Your repository must behave correctly even if connectivity from the cluster
// is disrupted. Reads and writes may fail in this case, but they must not
// return incorrect results.
//
// IMPORTANT: An analysis writes a substantial amount of data to your repository
// and then reads it back again.
// This consumes bandwidth on the network between the cluster and the
// repository, and storage space and I/O bandwidth on the repository itself.
// You must ensure this load does not affect other users of these systems.
// Analyses respect the repository settings `max_snapshot_bytes_per_sec` and
// `max_restore_bytes_per_sec` if available and the cluster setting
// `indices.recovery.max_bytes_per_sec` which you can use to limit the bandwidth
// they consume.
//
// NOTE: This API is intended for exploratory use by humans. You should expect
// the request parameters and the response format to vary in future versions.
//
// NOTE: Different versions of Elasticsearch may perform different checks for
// repository compatibility, with newer versions typically being stricter than
// older ones.
// A storage system that passes repository analysis with one version of
// Elasticsearch may fail with a different version.
// This indicates it behaves incorrectly in ways that the former version did not
// detect.
// You must work with the supplier of your storage system to address the
// incompatibilities detected by the repository analysis API in any version of
// Elasticsearch.
//
// NOTE: This API may not work correctly in a mixed-version cluster.
//
// *Implementation details*
//
// NOTE: This section of documentation describes how the repository analysis API
// works in this version of Elasticsearch, but you should expect the
// implementation to vary between versions. The request parameters and response
// format depend on details of the implementation so may also be different in
// newer versions.
//
// The analysis comprises a number of blob-level tasks, as set by the
// `blob_count` parameter and a number of compare-and-exchange operations on
// linearizable registers, as set by the `register_operation_count` parameter.
// These tasks are distributed over the data and master-eligible nodes in the
// cluster for execution.
//
// For most blob-level tasks, the executing node first writes a blob to the
// repository and then instructs some of the other nodes in the cluster to
// attempt to read the data it just wrote.
// The size of the blob is chosen randomly, according to the `max_blob_size` and
// `max_total_data_size` parameters.
// If any of these reads fails then the repository does not implement the
// necessary read-after-write semantics that Elasticsearch requires.
//
// For some blob-level tasks, the executing node will instruct some of its peers
// to attempt to read the data before the writing process completes.
// These reads are permitted to fail, but must not return partial data.
// If any read returns partial data then the repository does not implement the
// necessary atomicity semantics that Elasticsearch requires.
//
// For some blob-level tasks, the executing node will overwrite the blob while
// its peers are reading it.
// In this case the data read may come from either the original or the
// overwritten blob, but the read operation must not return partial data or a
// mix of data from the two blobs.
// If any of these reads returns partial data or a mix of the two blobs then the
// repository does not implement the necessary atomicity semantics that
// Elasticsearch requires for overwrites.
//
// The executing node will use a variety of different methods to write the blob.
// For instance, where applicable, it will use both single-part and multi-part
// uploads.
// Similarly, the reading nodes will use a variety of different methods to read
// the data back again.
// For instance they may read the entire blob from start to end or may read only
// a subset of the data.
//
// For some blob-level tasks, the executing node will cancel the write before it
// is complete.
// In this case, it still instructs some of the other nodes in the cluster to
// attempt to read the blob but all of these reads must fail to find the blob.
//
// Linearizable registers are special blobs that Elasticsearch manipulates using
// an atomic compare-and-exchange operation.
// This operation ensures correct and strongly-consistent behavior even when the
// blob is accessed by multiple nodes at the same time.
// The detailed implementation of the compare-and-exchange operation on
// linearizable registers varies by repository type.
// Repository analysis verifies that that uncontended compare-and-exchange
// operations on a linearizable register blob always succeed.
// Repository analysis also verifies that contended operations either succeed or
// report the contention but do not return incorrect results.
// If an operation fails due to contention, Elasticsearch retries the operation
// until it succeeds.
// Most of the compare-and-exchange operations performed by repository analysis
// atomically increment a counter which is represented as an 8-byte blob.
// Some operations also verify the behavior on small blobs with sizes other than
// 8 bytes.
package repositoryanalyze

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

type RepositoryAnalyze struct {
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

// NewRepositoryAnalyze type alias for index.
type NewRepositoryAnalyze func(repository string) *RepositoryAnalyze

// NewRepositoryAnalyzeFunc returns a new instance of RepositoryAnalyze with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRepositoryAnalyzeFunc(tp elastictransport.Interface) NewRepositoryAnalyze {
	return func(repository string) *RepositoryAnalyze {
		n := New(tp)

		n._repository(repository)

		return n
	}
}

// Analyze a snapshot repository.
// Analyze the performance characteristics and any incorrect behaviour found in
// a repository.
//
// The response exposes implementation details of the analysis which may change
// from version to version.
// The response body format is therefore not considered stable and may be
// different in newer versions.
//
// There are a large number of third-party storage systems available, not all of
// which are suitable for use as a snapshot repository by Elasticsearch.
// Some storage systems behave incorrectly, or perform poorly, especially when
// accessed concurrently by multiple clients as the nodes of an Elasticsearch
// cluster do. This API performs a collection of read and write operations on
// your repository which are designed to detect incorrect behaviour and to
// measure the performance characteristics of your storage system.
//
// The default values for the parameters are deliberately low to reduce the
// impact of running an analysis inadvertently and to provide a sensible
// starting point for your investigations.
// Run your first analysis with the default parameter values to check for simple
// problems.
// If successful, run a sequence of increasingly large analyses until you
// encounter a failure or you reach a `blob_count` of at least `2000`, a
// `max_blob_size` of at least `2gb`, a `max_total_data_size` of at least `1tb`,
// and a `register_operation_count` of at least `100`.
// Always specify a generous timeout, possibly `1h` or longer, to allow time for
// each analysis to run to completion.
// Perform the analyses using a multi-node cluster of a similar size to your
// production cluster so that it can detect any problems that only arise when
// the repository is accessed by many nodes at once.
//
// If the analysis fails, Elasticsearch detected that your repository behaved
// unexpectedly.
// This usually means you are using a third-party storage system with an
// incorrect or incompatible implementation of the API it claims to support.
// If so, this storage system is not suitable for use as a snapshot repository.
// You will need to work with the supplier of your storage system to address the
// incompatibilities that Elasticsearch detects.
//
// If the analysis is successful, the API returns details of the testing
// process, optionally including how long each operation took.
// You can use this information to determine the performance of your storage
// system.
// If any operation fails or returns an incorrect result, the API returns an
// error.
// If the API returns an error, it may not have removed all the data it wrote to
// the repository.
// The error will indicate the location of any leftover data and this path is
// also recorded in the Elasticsearch logs.
// You should verify that this location has been cleaned up correctly.
// If there is still leftover data at the specified location, you should
// manually remove it.
//
// If the connection from your client to Elasticsearch is closed while the
// client is waiting for the result of the analysis, the test is cancelled.
// Some clients are configured to close their connection if no response is
// received within a certain timeout.
// An analysis takes a long time to complete so you might need to relax any such
// client-side timeouts.
// On cancellation the analysis attempts to clean up the data it was writing,
// but it may not be able to remove it all.
// The path to the leftover data is recorded in the Elasticsearch logs.
// You should verify that this location has been cleaned up correctly.
// If there is still leftover data at the specified location, you should
// manually remove it.
//
// If the analysis is successful then it detected no incorrect behaviour, but
// this does not mean that correct behaviour is guaranteed.
// The analysis attempts to detect common bugs but it does not offer 100%
// coverage.
// Additionally, it does not test the following:
//
// * Your repository must perform durable writes. Once a blob has been written
// it must remain in place until it is deleted, even after a power loss or
// similar disaster.
// * Your repository must not suffer from silent data corruption. Once a blob
// has been written, its contents must remain unchanged until it is deliberately
// modified or deleted.
// * Your repository must behave correctly even if connectivity from the cluster
// is disrupted. Reads and writes may fail in this case, but they must not
// return incorrect results.
//
// IMPORTANT: An analysis writes a substantial amount of data to your repository
// and then reads it back again.
// This consumes bandwidth on the network between the cluster and the
// repository, and storage space and I/O bandwidth on the repository itself.
// You must ensure this load does not affect other users of these systems.
// Analyses respect the repository settings `max_snapshot_bytes_per_sec` and
// `max_restore_bytes_per_sec` if available and the cluster setting
// `indices.recovery.max_bytes_per_sec` which you can use to limit the bandwidth
// they consume.
//
// NOTE: This API is intended for exploratory use by humans. You should expect
// the request parameters and the response format to vary in future versions.
//
// NOTE: Different versions of Elasticsearch may perform different checks for
// repository compatibility, with newer versions typically being stricter than
// older ones.
// A storage system that passes repository analysis with one version of
// Elasticsearch may fail with a different version.
// This indicates it behaves incorrectly in ways that the former version did not
// detect.
// You must work with the supplier of your storage system to address the
// incompatibilities detected by the repository analysis API in any version of
// Elasticsearch.
//
// NOTE: This API may not work correctly in a mixed-version cluster.
//
// *Implementation details*
//
// NOTE: This section of documentation describes how the repository analysis API
// works in this version of Elasticsearch, but you should expect the
// implementation to vary between versions. The request parameters and response
// format depend on details of the implementation so may also be different in
// newer versions.
//
// The analysis comprises a number of blob-level tasks, as set by the
// `blob_count` parameter and a number of compare-and-exchange operations on
// linearizable registers, as set by the `register_operation_count` parameter.
// These tasks are distributed over the data and master-eligible nodes in the
// cluster for execution.
//
// For most blob-level tasks, the executing node first writes a blob to the
// repository and then instructs some of the other nodes in the cluster to
// attempt to read the data it just wrote.
// The size of the blob is chosen randomly, according to the `max_blob_size` and
// `max_total_data_size` parameters.
// If any of these reads fails then the repository does not implement the
// necessary read-after-write semantics that Elasticsearch requires.
//
// For some blob-level tasks, the executing node will instruct some of its peers
// to attempt to read the data before the writing process completes.
// These reads are permitted to fail, but must not return partial data.
// If any read returns partial data then the repository does not implement the
// necessary atomicity semantics that Elasticsearch requires.
//
// For some blob-level tasks, the executing node will overwrite the blob while
// its peers are reading it.
// In this case the data read may come from either the original or the
// overwritten blob, but the read operation must not return partial data or a
// mix of data from the two blobs.
// If any of these reads returns partial data or a mix of the two blobs then the
// repository does not implement the necessary atomicity semantics that
// Elasticsearch requires for overwrites.
//
// The executing node will use a variety of different methods to write the blob.
// For instance, where applicable, it will use both single-part and multi-part
// uploads.
// Similarly, the reading nodes will use a variety of different methods to read
// the data back again.
// For instance they may read the entire blob from start to end or may read only
// a subset of the data.
//
// For some blob-level tasks, the executing node will cancel the write before it
// is complete.
// In this case, it still instructs some of the other nodes in the cluster to
// attempt to read the blob but all of these reads must fail to find the blob.
//
// Linearizable registers are special blobs that Elasticsearch manipulates using
// an atomic compare-and-exchange operation.
// This operation ensures correct and strongly-consistent behavior even when the
// blob is accessed by multiple nodes at the same time.
// The detailed implementation of the compare-and-exchange operation on
// linearizable registers varies by repository type.
// Repository analysis verifies that that uncontended compare-and-exchange
// operations on a linearizable register blob always succeed.
// Repository analysis also verifies that contended operations either succeed or
// report the contention but do not return incorrect results.
// If an operation fails due to contention, Elasticsearch retries the operation
// until it succeeds.
// Most of the compare-and-exchange operations performed by repository analysis
// atomically increment a counter which is represented as an 8-byte blob.
// Some operations also verify the behavior on small blobs with sizes other than
// 8 bytes.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/repo-analysis-api.html
func New(tp elastictransport.Interface) *RepositoryAnalyze {
	r := &RepositoryAnalyze{
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
func (r *RepositoryAnalyze) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("_analyze")

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
func (r RepositoryAnalyze) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.repository_analyze")
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
		instrument.BeforeRequest(req, "snapshot.repository_analyze")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.repository_analyze", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.repository_analyze")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the RepositoryAnalyze query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a repositoryanalyze.Response
func (r RepositoryAnalyze) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.repository_analyze")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r RepositoryAnalyze) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.repository_analyze")
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
		err := fmt.Errorf("an error happened during the RepositoryAnalyze query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the RepositoryAnalyze headers map.
func (r *RepositoryAnalyze) Header(key, value string) *RepositoryAnalyze {
	r.headers.Set(key, value)

	return r
}

// Repository The name of the repository.
// API Name: repository
func (r *RepositoryAnalyze) _repository(repository string) *RepositoryAnalyze {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// BlobCount The total number of blobs to write to the repository during the test.
// For realistic experiments, you should set it to at least `2000`.
// API name: blob_count
func (r *RepositoryAnalyze) BlobCount(blobcount int) *RepositoryAnalyze {
	r.values.Set("blob_count", strconv.Itoa(blobcount))

	return r
}

// Concurrency The number of operations to run concurrently during the test.
// API name: concurrency
func (r *RepositoryAnalyze) Concurrency(concurrency int) *RepositoryAnalyze {
	r.values.Set("concurrency", strconv.Itoa(concurrency))

	return r
}

// Detailed Indicates whether to return detailed results, including timing information
// for every operation performed during the analysis.
// If false, it returns only a summary of the analysis.
// API name: detailed
func (r *RepositoryAnalyze) Detailed(detailed bool) *RepositoryAnalyze {
	r.values.Set("detailed", strconv.FormatBool(detailed))

	return r
}

// EarlyReadNodeCount The number of nodes on which to perform an early read operation while writing
// each blob.
// Early read operations are only rarely performed.
// API name: early_read_node_count
func (r *RepositoryAnalyze) EarlyReadNodeCount(earlyreadnodecount int) *RepositoryAnalyze {
	r.values.Set("early_read_node_count", strconv.Itoa(earlyreadnodecount))

	return r
}

// MaxBlobSize The maximum size of a blob to be written during the test.
// For realistic experiments, you should set it to at least `2gb`.
// API name: max_blob_size
func (r *RepositoryAnalyze) MaxBlobSize(bytesize string) *RepositoryAnalyze {
	r.values.Set("max_blob_size", bytesize)

	return r
}

// MaxTotalDataSize An upper limit on the total size of all the blobs written during the test.
// For realistic experiments, you should set it to at least `1tb`.
// API name: max_total_data_size
func (r *RepositoryAnalyze) MaxTotalDataSize(bytesize string) *RepositoryAnalyze {
	r.values.Set("max_total_data_size", bytesize)

	return r
}

// RareActionProbability The probability of performing a rare action such as an early read, an
// overwrite, or an aborted write on each blob.
// API name: rare_action_probability
func (r *RepositoryAnalyze) RareActionProbability(rareactionprobability string) *RepositoryAnalyze {
	r.values.Set("rare_action_probability", rareactionprobability)

	return r
}

// RarelyAbortWrites Indicates whether to rarely cancel writes before they complete.
// API name: rarely_abort_writes
func (r *RepositoryAnalyze) RarelyAbortWrites(rarelyabortwrites bool) *RepositoryAnalyze {
	r.values.Set("rarely_abort_writes", strconv.FormatBool(rarelyabortwrites))

	return r
}

// ReadNodeCount The number of nodes on which to read a blob after writing.
// API name: read_node_count
func (r *RepositoryAnalyze) ReadNodeCount(readnodecount int) *RepositoryAnalyze {
	r.values.Set("read_node_count", strconv.Itoa(readnodecount))

	return r
}

// RegisterOperationCount The minimum number of linearizable register operations to perform in total.
// For realistic experiments, you should set it to at least `100`.
// API name: register_operation_count
func (r *RepositoryAnalyze) RegisterOperationCount(registeroperationcount int) *RepositoryAnalyze {
	r.values.Set("register_operation_count", strconv.Itoa(registeroperationcount))

	return r
}

// Seed The seed for the pseudo-random number generator used to generate the list of
// operations performed during the test.
// To repeat the same set of operations in multiple experiments, use the same
// seed in each experiment.
// Note that the operations are performed concurrently so might not always
// happen in the same order on each run.
// API name: seed
func (r *RepositoryAnalyze) Seed(seed int) *RepositoryAnalyze {
	r.values.Set("seed", strconv.Itoa(seed))

	return r
}

// Timeout The period of time to wait for the test to complete.
// If no response is received before the timeout expires, the test is cancelled
// and returns an error.
// API name: timeout
func (r *RepositoryAnalyze) Timeout(duration string) *RepositoryAnalyze {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *RepositoryAnalyze) ErrorTrace(errortrace bool) *RepositoryAnalyze {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *RepositoryAnalyze) FilterPath(filterpaths ...string) *RepositoryAnalyze {
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
func (r *RepositoryAnalyze) Human(human bool) *RepositoryAnalyze {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *RepositoryAnalyze) Pretty(pretty bool) *RepositoryAnalyze {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
