package xkcdsearch

import (
	"strings"
)

const baseURL = "https://xkcd.com"

// Document wraps an xkcd.com comic.
//
type Document struct {
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	Published string `json:"published"`

	Title      string `json:"title"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	Link       string `json:"link,omitempty"`
	News       string `json:"news,omitempty"`
}

// URL returns the full URL to the comic.
//
func (d *Document) URL() string {
	var b strings.Builder

	b.WriteString(baseURL)
	b.WriteString("/")
	b.WriteString(d.ID)
	b.WriteString("/")

	return b.String()
}
