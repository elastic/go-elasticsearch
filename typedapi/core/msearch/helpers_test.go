package msearch

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestMsearch_AddSearch(t *testing.T) {
	type args struct {
		header types.MultisearchHeader
		body   types.MultisearchBody
	}
	tests := []struct {
		name    string
		args    args
		expects string
		wantErr bool
	}{
		{
			name: "nominal test case",
			args: args{
				header: types.MultisearchHeader{Index: []string{"foo"}},
				body:   types.MultisearchBody{Query: &types.Query{Term: map[string]types.TermQuery{"bar": {Value: "baz"}}}},
			},
			expects: `{"index":["foo"]}
{"query":{"term":{"bar":{"value":"baz"}}}}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.AddSearch(tt.args.header, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("AddSearch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !strings.EqualFold(r.buf.String(), tt.expects) {
				t.Errorf("wrong body")
			}
		})
	}
}
