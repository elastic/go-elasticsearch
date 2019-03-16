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
