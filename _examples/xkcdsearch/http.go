// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package xkcdsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

var httpLog zerolog.Logger

func init() {
	httpLog = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

}

// ServeHTTP is a handler returning search results in JSON.
//
func (s *Store) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	after := r.URL.Query().Get("a")
	if len(r.URL.RawQuery) > 0 {
		r.URL.RawQuery = fmt.Sprintf("?%s", r.URL.RawQuery)
	}
	httpLog.Info().Msg(fmt.Sprintf("%s %s%s", r.Method, r.URL.Path, r.URL.RawQuery))

	results, err := s.Search(query, after)
	if err != nil {
		httpLog.Error().Err(err).Msg("Failed to search the index")
		http.Error(w, fmt.Sprintf(`{ "error" : %q }`, err), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(results)
	if err != nil {
		httpLog.Error().Err(err).Msg("Failed to encode results to JSON")
		http.Error(w, fmt.Sprintf(`{ "error" : %q }`, err), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, bytes.NewReader(out))
	if err != nil {
		httpLog.Error().Err(err).Msg("Failed to write to response")
		http.Error(w, fmt.Sprintf(`{ "error" : %q }`, err), http.StatusInternalServerError)
		return
	}
}
