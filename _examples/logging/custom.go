// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

// This examples demonstrates how to implement the "estransport.Logger" interface with a custom type,
// and use it with the client for structured logging via the "rs/zerolog" package.

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/elastic/go-elasticsearch/v8"
)

// CustomLogger implements the estransport.Logger interface.
//
type CustomLogger struct {
	zerolog.Logger
}

// LogRoundTrip prints the information about request and response.
//
func (l *CustomLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	var (
		e    *zerolog.Event
		nReq int64
		nRes int64
	)

	// Set error level.
	//
	switch {
	case err != nil:
		e = l.Error()
	case res != nil && res.StatusCode > 0 && res.StatusCode < 300:
		e = l.Info()
	case res != nil && res.StatusCode > 299 && res.StatusCode < 500:
		e = l.Warn()
	case res != nil && res.StatusCode > 499:
		e = l.Error()
	default:
		e = l.Error()
	}

	// Count number of bytes in request and response.
	//
	if req != nil && req.Body != nil && req.Body != http.NoBody {
		nReq, _ = io.Copy(ioutil.Discard, req.Body)
	}
	if res != nil && res.Body != nil && res.Body != http.NoBody {
		nRes, _ = io.Copy(ioutil.Discard, res.Body)
	}

	// Log event.
	//
	e.Str("method", req.Method).
		Int("status_code", res.StatusCode).
		Dur("duration", dur).
		Int64("req_bytes", nReq).
		Int64("res_bytes", nRes).
		Msg(req.URL.String())

	return nil
}

// RequestBodyEnabled makes the client pass request body to logger
func (l *CustomLogger) RequestBodyEnabled() bool { return true }

// RequestBodyEnabled makes the client pass response body to logger
func (l *CustomLogger) ResponseBodyEnabled() bool { return true }

func main() {

	// ==============================================================================================
	//
	// Set up a logger
	//
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	// ==============================================================================================
	//
	// Pass the logger to the client
	//
	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Logger: &CustomLogger{log},
	})

	// ----------------------------------------------------------------------------------------------
	{
		es.Delete("test", "1")
		es.Exists("test", "1")
		es.Index("test", strings.NewReader(`{"title" : "logging"}`), es.Index.WithRefresh("true"))

		es.Search(
			es.Search.WithQuery("{FAIL"),
		)

		es.Search(
			es.Search.WithIndex("test"),
			es.Search.WithBody(strings.NewReader(`{"query" : {"match" : { "title" : "logging" } } }`)),
			es.Search.WithSize(1),
		)
	}
}
