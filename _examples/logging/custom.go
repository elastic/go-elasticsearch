// +build ignore

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/elastic/go-elasticsearch"
)

type CustomLogger struct {
	zerolog.Logger
}

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

func (l *CustomLogger) RequestBodyEnabled() bool  { return true }
func (l *CustomLogger) ResponseBodyEnabled() bool { return true }

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Logger: &CustomLogger{log},
	})

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
