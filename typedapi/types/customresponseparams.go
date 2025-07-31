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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"encoding/json"
)

// CustomResponseParams type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L987-L1125
type CustomResponseParams struct {
	// JsonParser Specifies the JSON parser that is used to parse the response from the custom
	// service.
	// Different task types require different json_parser parameters.
	// For example:
	// ```
	// # text_embedding
	// # For a response like this:
	//
	//	{
	//	 "object": "list",
	//	 "data": [
	//	     {
	//	       "object": "embedding",
	//	       "index": 0,
	//	       "embedding": [
	//	           0.014539449,
	//	           -0.015288644
	//	       ]
	//	     }
	//	 ],
	//	 "model": "text-embedding-ada-002-v2",
	//	 "usage": {
	//	     "prompt_tokens": 8,
	//	     "total_tokens": 8
	//	 }
	//	}
	//
	// # the json_parser definition should look like this:
	//
	//	"response":{
	//	  "json_parser":{
	//	    "text_embeddings":"$.data[*].embedding[*]"
	//	  }
	//	}
	//
	// # sparse_embedding
	// # For a response like this:
	//
	//	{
	//	  "request_id": "75C50B5B-E79E-4930-****-F48DBB392231",
	//	  "latency": 22,
	//	  "usage": {
	//	     "token_count": 11
	//	  },
	//	  "result": {
	//	     "sparse_embeddings": [
	//	        {
	//	          "index": 0,
	//	          "embedding": [
	//	            {
	//	              "token_id": 6,
	//	              "weight": 0.101
	//	            },
	//	            {
	//	              "token_id": 163040,
	//	              "weight": 0.28417
	//	            }
	//	          ]
	//	        }
	//	     ]
	//	  }
	//	}
	//
	// # the json_parser definition should look like this:
	//
	//	"response":{
	//	  "json_parser":{
	//	    "token_path":"$.result.sparse_embeddings[*].embedding[*].token_id",
	//	    "weight_path":"$.result.sparse_embeddings[*].embedding[*].weight"
	//	  }
	//	}
	//
	// # rerank
	// # For a response like this:
	//
	//	{
	//	  "results": [
	//	    {
	//	      "index": 3,
	//	      "relevance_score": 0.999071,
	//	      "document": "abc"
	//	    },
	//	    {
	//	      "index": 4,
	//	      "relevance_score": 0.7867867,
	//	      "document": "123"
	//	    },
	//	    {
	//	      "index": 0,
	//	      "relevance_score": 0.32713068,
	//	      "document": "super"
	//	    }
	//	  ],
	//	}
	//
	// # the json_parser definition should look like this:
	//
	//	"response":{
	//	  "json_parser":{
	//	    "reranked_index":"$.result.scores[*].index",    // optional
	//	    "relevance_score":"$.result.scores[*].score",
	//	    "document_text":"xxx"    // optional
	//	  }
	//	}
	//
	// # completion
	// # For a response like this:
	//
	//	{
	//	 "id": "chatcmpl-B9MBs8CjcvOU2jLn4n570S5qMJKcT",
	//	 "object": "chat.completion",
	//	 "created": 1741569952,
	//	 "model": "gpt-4.1-2025-04-14",
	//	 "choices": [
	//	   {
	//	    "index": 0,
	//	    "message": {
	//	      "role": "assistant",
	//	      "content": "Hello! How can I assist you today?",
	//	      "refusal": null,
	//	      "annotations": []
	//	    },
	//	    "logprobs": null,
	//	    "finish_reason": "stop"
	//	  }
	//	 ]
	//	}
	//
	// # the json_parser definition should look like this:
	//
	//	"response":{
	//	  "json_parser":{
	//	    "completion_result":"$.choices[*].message.content"
	//	  }
	//	}
	JsonParser json.RawMessage `json:"json_parser,omitempty"`
}

// NewCustomResponseParams returns a CustomResponseParams.
func NewCustomResponseParams() *CustomResponseParams {
	r := &CustomResponseParams{}

	return r
}
