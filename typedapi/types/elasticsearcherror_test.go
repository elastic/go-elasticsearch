package types

import (
	"reflect"
	"testing"
)

func TestElasticsearchError_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		data         string
		wantStatus   int
		wantType     string
		wantReason   *string
		wantMetadata map[string]string
		wantErr      bool
	}{
		{
			name:         "valid error and status",
			data:         `{"error": {"type": "some_type", "reason": "some_reason"}, "status": 404}`,
			wantStatus:   404,
			wantType:     "some_type",
			wantReason:   &[]string{"some_reason"}[0],
			wantMetadata: map[string]string{},
			wantErr:      false,
		},
		{
			name: "no shard failure",
			data: `{
					"error": {
                      "type": "exception",
                      "reason": "foo",
                      "caused_by": {
                        "type": "illegal_state_exception",
                        "reason": "bar"
                      }
                    }
				}`,
			wantStatus:   0,
			wantType:     "exception",
			wantReason:   &[]string{"foo"}[0],
			wantMetadata: map[string]string{},
			wantErr:      false,
		},
		{
			name:       "unknown field stored in metadata",
			data:       `{"error": {"type": "some_type"}, "status": 404, "foo": "bar"}`,
			wantStatus: 404,
			wantType:   "some_type",
			wantReason: nil,
			wantMetadata: map[string]string{
				"foo": `"bar"`,
			},
			wantErr: false,
		},
		{
			name:    "invalid json",
			data:    `{"error": {"type": "some_type", "reason": "some_reason", "status": 404`,
			wantErr: true,
		},
		{
			name:         "empty object",
			data:         `{}`,
			wantStatus:   0,
			wantType:     "",
			wantReason:   nil,
			wantMetadata: map[string]string{},
			wantErr:      false,
		},
		{
			name:    "non-object json",
			data:    `[]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e ElasticsearchError
			err := e.UnmarshalJSON([]byte(tt.data))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if e.Status != tt.wantStatus {
				t.Errorf("Status = %v, want %v", e.Status, tt.wantStatus)
			}
			if e.ErrorCause.Type != tt.wantType {
				t.Errorf("Type = %v, want %v", e.ErrorCause.Type, tt.wantType)
			}
			if (e.ErrorCause.Reason == nil) != (tt.wantReason == nil) ||
				(e.ErrorCause.Reason != nil && tt.wantReason != nil && *e.ErrorCause.Reason != *tt.wantReason) {
				t.Errorf("Reason = %v, want %v", e.ErrorCause.Reason, tt.wantReason)
			}
			gotMetadata := map[string]string{}
			for k, v := range e.ErrorCause.Metadata {
				gotMetadata[k] = string(v)
			}
			if !reflect.DeepEqual(gotMetadata, tt.wantMetadata) {
				t.Errorf("Metadata = %v, want %v", gotMetadata, tt.wantMetadata)
			}
		})
	}
}
