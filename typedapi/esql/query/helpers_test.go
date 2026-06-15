// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package query

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type Doc struct {
	Author      string `json:"author,omitempty"`
	Name        string `json:"name,omitempty"`
	PageCount   int    `json:"page_count,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
}

var testPayload = `{
  "columns": [
    {
      "name": "author",
      "type": "text"
    },
    {
      "name": "author.keyword",
      "type": "keyword"
    },
    {
      "name": "name",
      "type": "text"
    },
    {
      "name": "name.keyword",
      "type": "keyword"
    },
    {
      "name": "page_count",
      "type": "long"
    },
    {
      "name": "release_date",
      "type": "date"
    }
  ],
  "values": [
    [
      "James S.A. Corey",
      "James S.A. Corey",
      "Leviathan Wakes",
      "Leviathan Wakes",
      561,
      "2011-06-02T00:00:00.000Z"
    ],
    [
      "Dan Simmons",
      "Dan Simmons",
      "Hyperion",
      "Hyperion",
      482,
      "1989-05-26T00:00:00.000Z"
    ],
    [
      "Frank Herbert",
      "Frank Herbert",
      "Dune",
      "Dune",
      604,
      "1965-06-01T00:00:00.000Z"
    ],
    [
      "Frank Herbert",
      "Frank Herbert",
      "Dune Messiah",
      "Dune Messiah",
      331,
      "1969-10-15T00:00:00.000Z"
    ],
    [
      "Frank Herbert",
      "Frank Herbert",
      "Children of Dune",
      "Children of Dune",
      408,
      "1976-04-21T00:00:00.000Z"
    ],
    [
      "Frank Herbert",
      "Frank Herbert",
      "God Emperor of Dune",
      "God Emperor of Dune",
      454,
      "1981-05-28T00:00:00.000Z"
    ],
    [
      "Iain M. Banks",
      "Iain M. Banks",
      "Consider Phlebas",
      "Consider Phlebas",
      471,
      "1987-04-23T00:00:00.000Z"
    ],
    [
      "Peter F. Hamilton",
      "Peter F. Hamilton",
      "Pandora's Star",
      "Pandora's Star",
      768,
      "2004-03-02T00:00:00.000Z"
    ],
    [
      "Alastair Reynolds",
      "Alastair Reynolds",
      "Revelation Space",
      "Revelation Space",
      585,
      "2000-03-15T00:00:00.000Z"
    ],
    [
      "Vernor Vinge",
      "Vernor Vinge",
      "A Fire Upon the Deep",
      "A Fire Upon the Deep",
      613,
      "1992-06-01T00:00:00.000Z"
    ],
    [
      "Orson Scott Card",
      "Orson Scott Card",
      "Ender's Game",
      "Ender's Game",
      324,
      "1985-06-01T00:00:00.000Z"
    ],
    [
      "George Orwell",
      "George Orwell",
      "1984",
      "1984",
      328,
      "1985-06-01T00:00:00.000Z"
    ],
    [
      "Ray Bradbury",
      "Ray Bradbury",
      "Fahrenheit 451",
      "Fahrenheit 451",
      227,
      "1953-10-15T00:00:00.000Z"
    ],
    [
      "Aldous Huxley",
      "Aldous Huxley",
      "Brave New World",
      "Brave New World",
      268,
      "1932-06-01T00:00:00.000Z"
    ],
    [
      "Isaac Asimov",
      "Isaac Asimov",
      "Foundation",
      "Foundation",
      224,
      "1951-06-01T00:00:00.000Z"
    ],
    [
      "Lois Lowry",
      "Lois Lowry",
      "The Giver",
      "The Giver",
      208,
      "1993-04-26T00:00:00.000Z"
    ],
    [
      "Kurt Vonnegut",
      "Kurt Vonnegut",
      "Slaughterhouse-Five",
      "Slaughterhouse-Five",
      275,
      "1969-06-01T00:00:00.000Z"
    ],
    [
      "Douglas Adams",
      "Douglas Adams",
      "The Hitchhiker's Guide to the Galaxy",
      "The Hitchhiker's Guide to the Galaxy",
      180,
      "1979-10-12T00:00:00.000Z"
    ],
    [
      "Neal Stephenson",
      "Neal Stephenson",
      "Snow Crash",
      "Snow Crash",
      470,
      "1992-06-01T00:00:00.000Z"
    ],
    [
      "William Gibson",
      "William Gibson",
      "Neuromancer",
      "Neuromancer",
      271,
      "1984-07-01T00:00:00.000Z"
    ],
    [
      "Margaret Atwood",
      "Margaret Atwood",
      "The Handmaid's Tale",
      "The Handmaid's Tale",
      311,
      "1985-06-01T00:00:00.000Z"
    ],
    [
      "Robert A. Heinlein",
      "Robert A. Heinlein",
      "Starship Troopers",
      "Starship Troopers",
      335,
      "1959-12-01T00:00:00.000Z"
    ],
    [
      "Ursula K. Le Guin",
      "Ursula K. Le Guin",
      "The Left Hand of Darkness",
      "The Left Hand of Darkness",
      304,
      "1969-06-01T00:00:00.000Z"
    ],
    [
      "Robert A. Heinlein",
      "Robert A. Heinlein",
      "The Moon is a Harsh Mistress",
      "The Moon is a Harsh Mistress",
      288,
      "1966-04-01T00:00:00.000Z"
    ]
  ]
}`

type mockTransp struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

var successfullRoundTripFunc = func(*http.Request) (*http.Response, error) {
	res := &http.Response{}
	res.Header = http.Header{}
	res.Header.Add("Content-Type", "application/json")
	res.Body = io.NopCloser(strings.NewReader(testPayload))

	return res, nil
}

var badPayloadRoundTripFunc = func(*http.Request) (*http.Response, error) {
	res := &http.Response{}
	res.Header = http.Header{}
	res.Header.Add("Content-Type", "application/json")
	res.Body = io.NopCloser(strings.NewReader(`{ "columns":`))

	return res, nil
}

var customErrorRoundTripFunc = func(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("something really bad happened")
}

func (t mockTransp) Perform(request *http.Request) (*http.Response, error) {
	if t.RoundTripFunc == nil {
		return successfullRoundTripFunc(request)
	}
	return t.RoundTripFunc(request)
}

func TestHelper(t *testing.T) {
	type args struct {
		ctx       context.Context
		esqlQuery *Query
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    []T
		wantErr bool
	}
	tests := []testCase[Doc]{
		{
			name: "Simple deserialization",
			args: args{
				ctx: context.Background(),
				esqlQuery: &Query{
					transport: mockTransp{},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			want: []Doc{
				{Author: "James S.A. Corey", Name: "Leviathan Wakes", PageCount: 561, ReleaseDate: "2011-06-02T00:00:00.000Z"},
				{Author: "Dan Simmons", Name: "Hyperion", PageCount: 482, ReleaseDate: "1989-05-26T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Dune", PageCount: 604, ReleaseDate: "1965-06-01T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Dune Messiah", PageCount: 331, ReleaseDate: "1969-10-15T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Children of Dune", PageCount: 408, ReleaseDate: "1976-04-21T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "God Emperor of Dune", PageCount: 454, ReleaseDate: "1981-05-28T00:00:00.000Z"},
				{Author: "Iain M. Banks", Name: "Consider Phlebas", PageCount: 471, ReleaseDate: "1987-04-23T00:00:00.000Z"},
				{Author: "Peter F. Hamilton", Name: "Pandora's Star", PageCount: 768, ReleaseDate: "2004-03-02T00:00:00.000Z"},
				{Author: "Alastair Reynolds", Name: "Revelation Space", PageCount: 585, ReleaseDate: "2000-03-15T00:00:00.000Z"},
				{Author: "Vernor Vinge", Name: "A Fire Upon the Deep", PageCount: 613, ReleaseDate: "1992-06-01T00:00:00.000Z"},
				{Author: "Orson Scott Card", Name: "Ender's Game", PageCount: 324, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "George Orwell", Name: "1984", PageCount: 328, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "Ray Bradbury", Name: "Fahrenheit 451", PageCount: 227, ReleaseDate: "1953-10-15T00:00:00.000Z"},
				{Author: "Aldous Huxley", Name: "Brave New World", PageCount: 268, ReleaseDate: "1932-06-01T00:00:00.000Z"},
				{Author: "Isaac Asimov", Name: "Foundation", PageCount: 224, ReleaseDate: "1951-06-01T00:00:00.000Z"},
				{Author: "Lois Lowry", Name: "The Giver", PageCount: 208, ReleaseDate: "1993-04-26T00:00:00.000Z"},
				{Author: "Kurt Vonnegut", Name: "Slaughterhouse-Five", PageCount: 275, ReleaseDate: "1969-06-01T00:00:00.000Z"},
				{Author: "Douglas Adams", Name: "The Hitchhiker's Guide to the Galaxy", PageCount: 180, ReleaseDate: "1979-10-12T00:00:00.000Z"},
				{Author: "Neal Stephenson", Name: "Snow Crash", PageCount: 470, ReleaseDate: "1992-06-01T00:00:00.000Z"},
				{Author: "William Gibson", Name: "Neuromancer", PageCount: 271, ReleaseDate: "1984-07-01T00:00:00.000Z"},
				{Author: "Margaret Atwood", Name: "The Handmaid's Tale", PageCount: 311, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "Robert A. Heinlein", Name: "Starship Troopers", PageCount: 335, ReleaseDate: "1959-12-01T00:00:00.000Z"},
				{Author: "Ursula K. Le Guin", Name: "The Left Hand of Darkness", PageCount: 304, ReleaseDate: "1969-06-01T00:00:00.000Z"},
				{Author: "Robert A. Heinlein", Name: "The Moon is a Harsh Mistress", PageCount: 288, ReleaseDate: "1966-04-01T00:00:00.000Z"},
			},
			wantErr: false,
		},
		{
			name: "helper failure on bad response",
			args: args{
				ctx: context.Background(),
				esqlQuery: &Query{
					transport: mockTransp{badPayloadRoundTripFunc},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "helper failure on error",
			args: args{
				ctx: context.Background(),
				esqlQuery: &Query{
					transport: mockTransp{customErrorRoundTripFunc},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Helper[Doc](tt.args.ctx, tt.args.esqlQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("Helper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Helper() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIteratorHelper(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *Query
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    []Doc
		wantErr bool
	}
	tests := []testCase[Doc]{
		{
			name: "Simple iterator",
			args: args{
				ctx: context.Background(),
				query: &Query{
					transport: mockTransp{},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			want: []Doc{
				{Author: "James S.A. Corey", Name: "Leviathan Wakes", PageCount: 561, ReleaseDate: "2011-06-02T00:00:00.000Z"},
				{Author: "Dan Simmons", Name: "Hyperion", PageCount: 482, ReleaseDate: "1989-05-26T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Dune", PageCount: 604, ReleaseDate: "1965-06-01T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Dune Messiah", PageCount: 331, ReleaseDate: "1969-10-15T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "Children of Dune", PageCount: 408, ReleaseDate: "1976-04-21T00:00:00.000Z"},
				{Author: "Frank Herbert", Name: "God Emperor of Dune", PageCount: 454, ReleaseDate: "1981-05-28T00:00:00.000Z"},
				{Author: "Iain M. Banks", Name: "Consider Phlebas", PageCount: 471, ReleaseDate: "1987-04-23T00:00:00.000Z"},
				{Author: "Peter F. Hamilton", Name: "Pandora's Star", PageCount: 768, ReleaseDate: "2004-03-02T00:00:00.000Z"},
				{Author: "Alastair Reynolds", Name: "Revelation Space", PageCount: 585, ReleaseDate: "2000-03-15T00:00:00.000Z"},
				{Author: "Vernor Vinge", Name: "A Fire Upon the Deep", PageCount: 613, ReleaseDate: "1992-06-01T00:00:00.000Z"},
				{Author: "Orson Scott Card", Name: "Ender's Game", PageCount: 324, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "George Orwell", Name: "1984", PageCount: 328, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "Ray Bradbury", Name: "Fahrenheit 451", PageCount: 227, ReleaseDate: "1953-10-15T00:00:00.000Z"},
				{Author: "Aldous Huxley", Name: "Brave New World", PageCount: 268, ReleaseDate: "1932-06-01T00:00:00.000Z"},
				{Author: "Isaac Asimov", Name: "Foundation", PageCount: 224, ReleaseDate: "1951-06-01T00:00:00.000Z"},
				{Author: "Lois Lowry", Name: "The Giver", PageCount: 208, ReleaseDate: "1993-04-26T00:00:00.000Z"},
				{Author: "Kurt Vonnegut", Name: "Slaughterhouse-Five", PageCount: 275, ReleaseDate: "1969-06-01T00:00:00.000Z"},
				{Author: "Douglas Adams", Name: "The Hitchhiker's Guide to the Galaxy", PageCount: 180, ReleaseDate: "1979-10-12T00:00:00.000Z"},
				{Author: "Neal Stephenson", Name: "Snow Crash", PageCount: 470, ReleaseDate: "1992-06-01T00:00:00.000Z"},
				{Author: "William Gibson", Name: "Neuromancer", PageCount: 271, ReleaseDate: "1984-07-01T00:00:00.000Z"},
				{Author: "Margaret Atwood", Name: "The Handmaid's Tale", PageCount: 311, ReleaseDate: "1985-06-01T00:00:00.000Z"},
				{Author: "Robert A. Heinlein", Name: "Starship Troopers", PageCount: 335, ReleaseDate: "1959-12-01T00:00:00.000Z"},
				{Author: "Ursula K. Le Guin", Name: "The Left Hand of Darkness", PageCount: 304, ReleaseDate: "1969-06-01T00:00:00.000Z"},
				{Author: "Robert A. Heinlein", Name: "The Moon is a Harsh Mistress", PageCount: 288, ReleaseDate: "1966-04-01T00:00:00.000Z"},
			},
			wantErr: false,
		},
		{
			name: "Iterator failure on bad response",
			args: args{
				ctx: context.Background(),
				query: &Query{
					transport: mockTransp{badPayloadRoundTripFunc},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Iterator failure on error",
			args: args{
				ctx: context.Background(),
				query: &Query{
					transport: mockTransp{customErrorRoundTripFunc},
					values:    make(url.Values),
					headers:   make(http.Header),
					buf:       bytes.NewBuffer(nil),
					req: &Request{
						Query: `FROM docs`,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter, err := NewIteratorHelper[Doc](tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIteratorHelper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == false {
				var i int
				for iter.More() {
					if doc, err := iter.Next(); err != nil {
						t.Errorf("Iterator.Next() error = %v", err)
					} else {
						if reflect.DeepEqual(doc, tt.want[i]) {
							t.Errorf("NewIteratorHelper() got = %v, want %v", doc, tt.want)
						}
					}
					i++
				}
			}
		})
	}
}
