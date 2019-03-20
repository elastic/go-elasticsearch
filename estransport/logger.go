package estransport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (l *Logger) logRoundTrip(req *http.Request, res *http.Response, dur time.Duration, err error) {
	switch l.format {
	case LogFormatText, LogFormatColor:
		l.writeRoundTripText(req, res, dur, err)
		l.writeRequestBodyText(req)
		l.writeResponseBodyText(res)
		l.writeErrorText(err)

	case LogFormatCurl:
		l.writeRoundTripCurl(req, res, dur, err)
	case LogFormatJSON:
		l.writeRoundTripJSON(req, res, dur, err)
	}
}

func (l *Logger) writeErrorText(err error) {
	if err != nil {
		fmt.Fprintf(l.output, "! ERROR: %v", err)
	}
}

func (l *Logger) writeRoundTripText(req *http.Request, res *http.Response, dur time.Duration, err error) {
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

func (l *Logger) writeRoundTripJSON(req *http.Request, res *http.Response, dur time.Duration, err error) {
	// https://github.com/elastic/ecs/blob/master/schemas/http.yml
	//
	// TODO(karmi): Research performance optimization of using sync.Pool
	var b bytes.Buffer

	b.WriteRune('{')
	// -- Timestamp
	b.WriteString(`"@timestamp":"`)
	b.WriteString(time.Now().UTC().Format(time.RFC3339))
	b.WriteRune('"')
	// -- Event
	b.WriteString(`,"event":{`)
	b.WriteString(`"duration":`)
	b.WriteString(strconv.FormatInt(dur.Nanoseconds(), 10))
	b.WriteRune('}')
	// -- HTTP
	b.WriteString(`,"http":`)
	// ---- Request
	b.WriteString(`{"request":{`)
	b.WriteString(`"scheme":`)
	b.WriteString(strconv.Quote(req.URL.Scheme))
	b.WriteString(`,"host":`)
	b.WriteString(strconv.Quote(req.URL.Host))
	b.WriteString(`,"method":`)
	b.WriteString(strconv.Quote(req.Method))
	b.WriteString(`,"path":`)
	b.WriteString(strconv.Quote(req.URL.Path))
	b.WriteString(`,"query":`)
	b.WriteString(strconv.Quote(req.URL.RawQuery))
	if req.Body != nil && req.Body != http.NoBody {
		var body bytes.Buffer
		body.ReadFrom(req.Body)
		b.WriteString(`,"body":{`)
		b.WriteString(`"content":`)
		b.WriteString(strconv.Quote(body.String()))
		b.WriteRune('}') // Close "http.request.body"
	}
	b.WriteRune('}') // Close "http.request"
	// ---- Response
	if res != nil {
		b.WriteString(`,"response":{`)
		b.WriteString(`"status_code":`)
		b.WriteString(strconv.Itoa(res.StatusCode))
		if res.Body != nil && res.Body != http.NoBody {
			var body bytes.Buffer
			body.ReadFrom(res.Body)
			b.WriteString(`,"body":{`)
			b.WriteString(`"content":`)
			b.WriteString(strconv.Quote(body.String()))
			b.WriteRune('}') // Close "http.response.body"
		}
		b.WriteRune('}') // Close "http.response"
	}
	b.WriteRune('}') // Close "http"
	// -- Error
	if err != nil {
		b.WriteString(`,"error":{"message":`)
		b.WriteString(strconv.Quote(err.Error()))
		b.WriteRune('}') // Close "error"
	}
	b.WriteRune('}')
	b.WriteRune('\n')
	b.WriteTo(l.output)
}

func (l *Logger) writeRoundTripCurl(req *http.Request, res *http.Response, dur time.Duration, err error) {
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
