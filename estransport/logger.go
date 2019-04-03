package estransport

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Logger defines a logger.
//
type Logger interface {
	LogRoundTrip(*http.Request, *http.Response, error, time.Time, time.Duration)
	IsLoggingRequestBody() bool
	IsLoggingResponseBody() bool
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

// logger represents the default logger.
//
type logger struct {
	output io.Writer
	format LogFormat

	isLoggingRequestBody  bool
	isLoggingResponseBody bool
}

// newLogger returns new logger, when w is not nil, otherwise it returns nil.
//
func newLogger(w io.Writer, f LogFormat, reqBody bool, resBody bool) Logger {
	if w == nil {
		return nil
	}

	if f == LogFormatNone {
		f = LogFormatText
	}
	return &logger{output: w, format: f, isLoggingRequestBody: reqBody, isLoggingResponseBody: resBody}
}

func (l *logger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration) {
	switch l.format {
	case LogFormatText:
		l.writeRoundTripText(req, res, err, start, dur)
	case LogFormatColor:
		l.writeRoundTripColor(req, res, err, start, dur)
	case LogFormatCurl:
		l.writeRoundTripCurl(req, res, err, start, dur)
	case LogFormatJSON:
		l.writeRoundTripJSON(req, res, err, start, dur)
	}
}

// IsLoggingRequestBody returns true when the request body is logged.
//
func (l *logger) IsLoggingRequestBody() bool {
	return l.isLoggingRequestBody
}

// IsLoggingResponseBody returns true when the response body is logged.
//
func (l *logger) IsLoggingResponseBody() bool {
	return l.isLoggingResponseBody
}

func (l *logger) writeRoundTripText(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) {
	fmt.Fprintf(l.output, "%s %s %s [status:%d request:%s]\n",
		start.Format(time.RFC3339),
		req.Method,
		req.URL.String(),
		l.resStatusCode(res),
		dur.Truncate(time.Millisecond),
	)
	if l.IsLoggingRequestBody() && req != nil && req.Body != nil && req.Body != http.NoBody {
		l.writeRequestBodyText(req, ">")
	}
	if l.IsLoggingResponseBody() && res != nil && res.Body != nil && res.Body != http.NoBody {
		l.writeResponseBodyText(res, "<")
	}
	l.writeErrorText(err)
}

func (l *logger) writeRequestBodyText(req *http.Request, prefix string) {
	b1, b2, err := l.duplicateBody(req.Body)
	if err != nil {
		return
	}
	req.Body = ioutil.NopCloser(b2)

	scanner := bufio.NewScanner(b1)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			fmt.Fprintf(l.output, "%s %s\n", prefix, s)
		}
	}
}

func (l *logger) writeResponseBodyText(res *http.Response, prefix string) {
	b1, b2, err := l.duplicateBody(res.Body)
	if err != nil {
		return
	}
	res.Body = ioutil.NopCloser(b2)

	scanner := bufio.NewScanner(b1)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			fmt.Fprintf(l.output, "%s %s\n", prefix, s)
		}
	}
}

func (l *logger) writeErrorText(err error) {
	if err != nil {
		fmt.Fprintf(l.output, "! ERROR: %v\n", err)
	}
}

func (l *logger) writeRoundTripColor(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) {
	query, _ := url.QueryUnescape(req.URL.RawQuery)
	if query != "" {
		query = "?" + query
	}

	var (
		status string
		color  string
	)

	if res != nil {
		status = res.Status
		switch {
		case res.StatusCode > 0 && res.StatusCode < 300:
			color = "\x1b[32m"
		case res.StatusCode > 299 && res.StatusCode < 500:
			color = "\x1b[33m"
		case res.StatusCode > 499:
			color = "\x1b[31m"
		}
	} else {
		status = "ERROR"
		color = "\x1b[31;4m"
	}

	fmt.Fprintf(l.output, "%6s \x1b[1;4m%s://%s%s\x1b[0m%s %s%s\x1b[0m \x1b[2m%s\x1b[0m\n",
		req.Method,
		req.URL.Scheme,
		req.URL.Host,
		req.URL.Path,
		query,
		color,
		status,
		dur.Truncate(time.Millisecond),
	)

	if l.IsLoggingRequestBody() && req != nil && req.Body != nil && req.Body != http.NoBody {
		fmt.Fprint(l.output, "\x1b[2m")
		l.writeRequestBodyText(req, "       »")
		fmt.Fprint(l.output, "\x1b[0m")
	}

	if l.IsLoggingResponseBody() && res != nil && res.Body != nil && res.Body != http.NoBody {
		fmt.Fprint(l.output, "\x1b[2m")
		l.writeResponseBodyText(res, "       «")
		fmt.Fprint(l.output, "\x1b[0m")
	}

	if err != nil {
		fmt.Fprintf(l.output, "\x1b[31;1m» ERROR \x1b[31m%v\x1b[0m\n", err)
	}

	if l.IsLoggingRequestBody() || l.IsLoggingResponseBody() {
		fmt.Fprintf(l.output, "\x1b[2m%s\x1b[0m\n", strings.Repeat("─", 80))
	}
}

func (l *logger) writeRoundTripCurl(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) {
	var b bytes.Buffer

	var query string
	qvalues := url.Values{}
	for k, v := range req.URL.Query() {
		if k == "pretty" {
			continue
		}
		for _, qv := range v {
			qvalues.Add(k, qv)
		}
	}
	if len(qvalues) > 0 {
		query = qvalues.Encode()
	}

	b.WriteString(`curl`)
	if req.Method == "HEAD" {
		b.WriteString(" --head")
	} else {
		fmt.Fprintf(&b, " -X %s", req.Method)
	}

	if len(req.Header) > 0 {
		for k, vv := range req.Header {
			if k == "Authorization" {
				continue
			}
			v := strings.Join(vv, ",")
			b.WriteString(fmt.Sprintf(" -H '%s: %s'", k, v))
		}
	}

	b.WriteString(" 'http://localhost:9200")
	b.WriteString(req.URL.Path)
	b.WriteString("?pretty")
	if query != "" {
		fmt.Fprintf(&b, "&%s", query)
	}
	b.WriteString("'")

	if req.Body != nil && req.Body != http.NoBody {
		b1, b2, err := l.duplicateBody(req.Body)
		if err != nil {
			return
		}
		req.Body = ioutil.NopCloser(b2)

		b.Grow(b1.Len())
		b.WriteString(" -d \\\n'")
		json.Indent(&b, b1.Bytes(), "", " ")
		b.WriteString("'")
	}

	b.WriteRune('\n')

	var status string
	if res != nil {
		status = res.Status
	} else {
		status = "N/A"
	}
	fmt.Fprintf(&b, "# => %s [%s] %s\n", start.UTC().Format(time.RFC3339), status, dur.Truncate(time.Millisecond))
	if l.IsLoggingResponseBody() && res != nil && res.Body != nil && res.Body != http.NoBody {
		b1, b2, err := l.duplicateBody(res.Body)
		if err != nil {
			return
		}
		res.Body = ioutil.NopCloser(b2)

		b.Grow(b1.Len())
		b.WriteString("# ")
		json.Indent(&b, b1.Bytes(), "# ", " ")
	}

	b.WriteString("\n")
	if l.IsLoggingResponseBody() && res != nil && res.Body != nil && res.Body != http.NoBody {
		b.WriteString("\n")
	}

	b.WriteTo(l.output)
}

func (l *logger) writeRoundTripJSON(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) {
	// https://github.com/elastic/ecs/blob/master/schemas/http.yml
	//
	// TODO(karmi): Research performance optimization of using sync.Pool

	bsize := 200
	var b = bytes.NewBuffer(make([]byte, 0, bsize))
	var v = make([]byte, 0, bsize)

	appendTime := func(t time.Time) {
		v = v[:0]
		v = t.AppendFormat(v, time.RFC3339)
		b.Write(v)
	}

	appendQuote := func(s string) {
		v = v[:0]
		v = strconv.AppendQuote(v, s)
		b.Write(v)
	}

	appendInt := func(i int64) {
		v = v[:0]
		v = strconv.AppendInt(v, i, 10)
		b.Write(v)
	}

	port := req.URL.Port()

	b.WriteRune('{')
	// -- Timestamp
	b.WriteString(`"@timestamp":"`)
	appendTime(start.UTC())
	b.WriteRune('"')
	// -- Event
	b.WriteString(`,"event":{`)
	b.WriteString(`"duration":`)
	appendInt(dur.Nanoseconds())
	b.WriteRune('}')
	// -- URL
	b.WriteString(`,"url":{`)
	b.WriteString(`"scheme":`)
	appendQuote(req.URL.Scheme)
	b.WriteString(`,"domain":`)
	appendQuote(req.URL.Hostname())
	if port != "" {
		b.WriteString(`,"port":`)
		b.WriteString(port)
	}
	b.WriteString(`,"path":`)
	appendQuote(req.URL.Path)
	b.WriteString(`,"query":`)
	appendQuote(req.URL.RawQuery)
	b.WriteRune('}') // Close "url"
	// -- HTTP
	b.WriteString(`,"http":`)
	// ---- Request
	b.WriteString(`{"request":{`)
	b.WriteString(`"method":`)
	appendQuote(req.Method)
	if l.IsLoggingRequestBody() && req != nil && req.Body != nil && req.Body != http.NoBody {
		b.WriteString(`,"body":`)
		b1, b2, err := l.duplicateBody(req.Body)
		if err != nil {
			return
		}
		req.Body = ioutil.NopCloser(b2)

		b.Grow(b1.Len())
		appendQuote(b1.String())
	}
	b.WriteRune('}') // Close "http.request"
	// ---- Response
	if res != nil {
		b.WriteString(`,"response":{`)
		b.WriteString(`"status_code":`)
		appendInt(int64(l.resStatusCode(res)))
		if l.IsLoggingResponseBody() && res != nil && res.Body != nil && res.Body != http.NoBody {
			b.WriteString(`,"body":`)
			b1, b2, err := l.duplicateBody(res.Body)
			if err != nil {
				return
			}
			res.Body = ioutil.NopCloser(b2)

			b.Grow(b1.Len())
			appendQuote(b1.String())
		}
		b.WriteRune('}') // Close "http.response"
	}
	b.WriteRune('}') // Close "http"
	// -- Error
	if err != nil {
		b.WriteString(`,"error":{"message":`)
		appendQuote(err.Error())
		b.WriteRune('}') // Close "error"
	}
	b.WriteRune('}')
	b.WriteRune('\n')
	b.WriteTo(l.output)
}

func (l *logger) duplicateBody(body io.ReadCloser) (*bytes.Buffer, *bytes.Buffer, error) {
	var (
		b1 bytes.Buffer
		b2 bytes.Buffer
		tr = io.TeeReader(body, &b2)
	)
	_, err := b1.ReadFrom(tr)
	if err != nil {
		return nil, nil, err
	}
	defer func() { body.Close() }()

	return &b1, &b2, nil
}

func (l *logger) resStatusCode(res *http.Response) int {
	if res != nil {
		return res.StatusCode
	}
	return -1
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
