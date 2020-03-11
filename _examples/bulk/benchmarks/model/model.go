// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

//go:generate easyjson $GOFILE

package model

import (
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

// BulkIndexerResponse wraps the esutil.BulkIndexerResponse,
// and implements the esutil.UnmarshalFromReader() method,
// in order to demonstrate usage of a third-party JSON decoder,
// such as "mailru/easyjson".
//
// easyjson:json
type BulkIndexerResponse struct {
	esutil.BulkIndexerResponse
}
