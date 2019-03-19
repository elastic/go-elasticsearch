package estransport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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

// Logger represents the default logger.
//
type Logger struct {
	output io.Writer
	format LogFormat
}

// NewLogger returns new logger, when w is not nil, otherwise it returns nil.
//
func NewLogger(w io.Writer, f LogFormat) *Logger {
	if w == nil {
		return nil
	}
	if f == LogFormatNone {
		f = LogFormatText
	}
	return &Logger{output: w, format: f}
}

func (l *Logger) logRoundTrip(req *http.Request, res *http.Response, dur time.Duration) {
	switch l.format {
	case LogFormatText, LogFormatColor:
		l.writeRoundTripText(req, res, dur)
		l.writeRequestBodyText(req)
		l.writeResponseBodyText(res)
	case LogFormatCurl:
		l.writeRoundTripCurl(req, res, dur)
	case LogFormatJSON:
		l.writeRoundTripJSON(req, res, dur)
	}
}

func (l *Logger) logError(err error) {
	fmt.Fprintf(l.output, "! ERROR: %v", err)
}

func (l *Logger) writeRoundTripText(req *http.Request, res *http.Response, dur time.Duration) {
	fmt.Fprintf(l.output, "%s %s %s [status:%d request:%s]\n",
		time.Now().Format(time.RFC3339),
		req.Method,
		// TODO(karmi): Unescape raw query
		req.URL.String(),
		res.StatusCode,
		dur.Truncate(time.Millisecond),
	)
}

func (l *Logger) writeRequestBodyText(req *http.Request) {
	if req.Body != nil && req.Body != http.NoBody {
		// TODO(karmi): Use bufio.Scan
		body, err := ioutil.ReadAll(req.Body)
		if err == nil {
			for _, line := range strings.Split(string(body), "\n") {
				if line != "" {
					fmt.Fprintf(l.output, "> %s\n", line)
				}
			}
		}
	}
}

func (l *Logger) writeResponseBodyText(res *http.Response) {
	if res.Body != nil && res.Body != http.NoBody {
		// TODO(karmi): Use bufio.Scan
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			defer func() { res.Body = ioutil.NopCloser(bytes.NewReader(body)) }()
			defer func() { res.Body.Close() }()
			for _, line := range strings.Split(string(body), "\n") {
				if line != "" {
					fmt.Fprintf(l.output, "< %s\n", line)
				}
			}
		}
	}
}

func (l *Logger) writeRoundTripJSON(req *http.Request, res *http.Response, dur time.Duration) {
	// https://github.com/elastic/ecs/blob/master/code/go/ecs/http.go
	// https://github.com/elastic/ecs/blob/master/schemas/http.yml
	fmt.Fprintln(l.output, "### TODO: JSON ###")
}

func (l *Logger) writeRoundTripCurl(req *http.Request, res *http.Response, dur time.Duration) {
	fmt.Fprintln(l.output, "### TODO: Curl ###")
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
