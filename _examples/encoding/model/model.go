// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

//go:generate easyjson -all -snake_case $GOFILE

package model

import "time"

type Article struct {
	ID        uint
	Title     string
	Body      string
	Published time.Time
	Author    *Author
}

type Author struct {
	FirstName string
	LastName  string
}
