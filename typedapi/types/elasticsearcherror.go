package types

import (
	"fmt"
	"strings"
)

type ElasticsearchError struct {
	ErrorCause ErrorCause `json:"error"`
	Status     int        `json:"status"`
}

// Error implements error string serialization of the ElasticsearchError.
func (e ElasticsearchError) Error() string {
	var reason string
	if e.ErrorCause.Reason != nil {
		reason = *e.ErrorCause.Reason
	}
	return fmt.Sprintf("status: %d, failed: [%s], reason: %s", e.Status, e.ErrorCause.Type, reason)
}

// Is implements errors.Is interface to allow value comparison within ElasticsearchError.
// It checks for always present values only: Status & ErrorCause.Type.
func (e ElasticsearchError) Is(err error) bool {
	prefix := fmt.Sprintf("status: %d, failed: [%s]", e.Status, e.ErrorCause.Type)
	if strings.HasPrefix(err.Error(), prefix) {
		return true
	}
	return false
}

// As implements errors.As interface to allow type matching of ElasticsearchError.
func (e ElasticsearchError) As(err interface{}) bool {
	if _, ok := err.(*ElasticsearchError); ok {
		return true
	}
	return false
}

// NewElasticsearchError returns a ElasticsearchError.
func NewElasticsearchError() *ElasticsearchError {
	r := &ElasticsearchError{}

	return r
}
