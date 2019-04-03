package estransport

import (
	"io"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/estransport/logger"
)

// Logger defines an interface for logging request and response.
//
type Logger interface {
	LogRoundTrip(*http.Request, *http.Response, error, time.Time, time.Duration) error
	RequestBodyEnabled() bool
	ResponseBodyEnabled() bool
}

// LogFormat defines the logger output format.
//
type LogFormat int

// Log formats:
const (
	LogFormatNone  LogFormat = iota
	LogFormatText            // Plain text
	LogFormatColor           // Terminal-optimized plain text
	LogFormatCurl            // Runnable curl command
	LogFormatJSON            // Structured output
)

// NewLogger returns new logger, when w is not nil, otherwise it returns nil.
//
func NewLogger(w io.Writer, f LogFormat, reqBody bool, resBody bool) Logger {
	if w == nil {
		return nil
	}

	var l Logger

	switch f {
	case LogFormatText:
		l = &logger.Text{Output: w, EnableRequestBody: reqBody, EnableResponseBody: resBody}
	case LogFormatColor:
		l = &logger.Color{Output: w, EnableRequestBody: reqBody, EnableResponseBody: resBody}
	case LogFormatCurl:
		l = &logger.Curl{Output: w, EnableRequestBody: reqBody, EnableResponseBody: resBody}
	case LogFormatJSON:
		l = &logger.JSON{Output: w, EnableRequestBody: reqBody, EnableResponseBody: resBody}
	default:
		l = &logger.Text{Output: w, EnableRequestBody: reqBody, EnableResponseBody: resBody}
	}

	return l
}

// String returns LogFormat as a string.
//
func (f LogFormat) String() string {
	switch f {
	case LogFormatText:
		return "text"
	case LogFormatColor:
		return "color"
	case LogFormatCurl:
		return "curl"
	case LogFormatJSON:
		return "json"
	}
	return "unknown"
}
