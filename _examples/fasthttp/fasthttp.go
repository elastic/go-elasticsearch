package fasthttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/valyala/fasthttp"
)

// Transport implements the estransport interface with
// the github.com/valyala/fasthttp HTTP client.
//
type Transport struct{}

// RoundTrip performs the request and returns a response or error
//
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	freq := fasthttp.AcquireRequest()
	fres := fasthttp.AcquireResponse()
	t.copyRequest(freq, req)
	err := fasthttp.Do(freq, fres)
	if err != nil {
		return nil, err
	}

	res := &http.Response{Header: make(http.Header)}
	t.copyResponse(res, fres)

	return res, nil
}

// copyRequest converts http.Request to fasthttp.Request
//
func (t *Transport) copyRequest(dst *fasthttp.Request, src *http.Request) *fasthttp.Request {
	dst.Reset()

	dst.SetHost(src.Host)
	dst.SetRequestURI(src.URL.String())

	dst.Header.SetRequestURI(src.URL.String())
	dst.Header.SetMethod(src.Method)

	for k, vv := range src.Header {
		for _, v := range vv {
			dst.Header.Set(k, v)
		}
	}

	if src.Body != nil {
		var b bytes.Buffer
		io.Copy(&b, src.Body)

		dst.SetBody(b.Bytes())
	}

	return dst
}

// copyResponse converts http.Response to fasthttp.Response
//
func (t *Transport) copyResponse(dst *http.Response, src *fasthttp.Response) *http.Response {
	dst.StatusCode = src.StatusCode()

	src.Header.VisitAll(func(k, v []byte) {
		dst.Header.Set(string(k), string(v))
	})

	if src.Body != nil {
		dst.Body = ioutil.NopCloser(bytes.NewReader(src.Body()))
	}

	return dst
}
