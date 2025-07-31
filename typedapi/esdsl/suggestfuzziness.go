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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _suggestFuzziness struct {
	v *types.SuggestFuzziness
}

func NewSuggestFuzziness() *_suggestFuzziness {

	return &_suggestFuzziness{v: types.NewSuggestFuzziness()}

}

func (s *_suggestFuzziness) Fuzziness(fuzziness types.FuzzinessVariant) *_suggestFuzziness {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

func (s *_suggestFuzziness) MinLength(minlength int) *_suggestFuzziness {

	s.v.MinLength = &minlength

	return s
}

func (s *_suggestFuzziness) PrefixLength(prefixlength int) *_suggestFuzziness {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_suggestFuzziness) Transpositions(transpositions bool) *_suggestFuzziness {

	s.v.Transpositions = &transpositions

	return s
}

func (s *_suggestFuzziness) UnicodeAware(unicodeaware bool) *_suggestFuzziness {

	s.v.UnicodeAware = &unicodeaware

	return s
}

func (s *_suggestFuzziness) SuggestFuzzinessCaster() *types.SuggestFuzziness {
	return s.v
}
