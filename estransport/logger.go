package estransport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func logRoundTrip(w io.Writer, req *http.Request, res *http.Response, dur time.Duration) {
	fmt.Fprintf(w, "%s %s %s [status:%d request:%s]\n",
		time.Now().Format(time.RFC3339),
		req.Method,
		req.URL.String(),
		res.StatusCode,
		dur.Truncate(time.Millisecond),
	)
	if req.Body != nil {
		// TODO(karmi): Capture the request body before performing the request
		fmt.Fprintln(w, "> TODO: Capture and print request body")
	}
}

func logResponseBody(w io.Writer, res *http.Response) {
	if res.Body != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			defer func() { res.Body = ioutil.NopCloser(bytes.NewReader(body)) }()
			defer func() { res.Body.Close() }()
			for _, line := range strings.Split(string(body), "\n") {
				if line != "" {
					log.Printf("< %s\n", line)
				}
			}
		}
	}
}

func logError(w io.Writer, err error) {
	fmt.Fprintf(w, "! ERROR: %v", err)
}
