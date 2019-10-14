// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch"
)

var (
	indexSetup     bool
	crawlerWorkers int
)

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().BoolVar(&indexSetup, "setup", false, "Create Elasticsearch index")
	indexCmd.Flags().IntVar(&crawlerWorkers, "workers", 25, "Number of concurrent workers")
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index xkcd.com into Elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		crawler := Crawler{
			log: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
				Level(func() zerolog.Level {
					if os.Getenv("DEBUG") != "" {
						return zerolog.DebugLevel
					} else {
						return zerolog.InfoLevel
					}
				}()).
				With().
				Timestamp().
				Logger(),

			workers: crawlerWorkers,
			queue:   make(chan string, crawlerWorkers),
			reURL:   regexp.MustCompile(`https\://xkcd\.com/(?P<ID>\d+)/.*\.json`),
			nextURL: func(c *Crawler, u string) string {
				id, err := c.documentIDFromURL(u)
				if err != nil || id <= 1 {
					return ""
				}
				return fmt.Sprintf("https://xkcd.com/%d/info.0.json", id-1)
			},

			StartURL: "https://xkcd.com/info.0.json",
		}

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			crawler.log.Fatal().Err(err).Msg("Error creating Elasticsearch client")
		}

		config := xkcdsearch.StoreConfig{Client: es, IndexName: IndexName}
		store, err := xkcdsearch.NewStore(config)
		if err != nil {
			crawler.log.Fatal().Err(err).Msg("Cannot create store")
		}
		crawler.store = store

		if indexSetup {
			crawler.log.Info().Msg("Creating index with mapping")
			if err := crawler.setupIndex(); err != nil {
				crawler.log.Fatal().Err(err).Msg("Cannot create Elasticsearch index")
			}
		}

		crawler.log.Info().Msgf("Starting the crawl with %d workers at <%s>", crawler.workers, crawler.StartURL)
		crawler.Run()
	},
}

// Crawler allows to crawl and index comics from xkcd.com.
//
type Crawler struct {
	store *xkcdsearch.Store
	log   zerolog.Logger

	workers int
	queue   chan string
	reURL   *regexp.Regexp
	nextURL func(crawler *Crawler, currentURL string) (nextURL string)

	StartURL string
}

// Run launches the crawler.
//
func (c *Crawler) Run() {
	var (
		wg      sync.WaitGroup
		nextURL string
	)

	rand.Seed(time.Now().Unix())

	// Setup worker goroutines
	for i := 1; i < c.workers+1; i++ {
		c.log.Debug().Msgf("Creating worker %d", i)
		wg.Add(1)
		go func() {
			defer wg.Done()

			for url := range c.queue {
				c.ProcessURL(url)
			}
		}()
	}

	// Start crawling
	if doc := c.ProcessURL(c.StartURL); doc.ID != "" {
		id, err := strconv.Atoi(doc.ID)
		if err != nil {
			c.log.Fatal().Err(err).Msg("Cannot get latest comic ID")
		}
		nextURL = fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
	}

	for {
		if nextURL = c.NextURL(nextURL); nextURL == "" {
			close(c.queue)
			break
		}

		c.log.Debug().Str("URL", nextURL).Msg("Adding URL to queue")
		c.queue <- nextURL
	}

	wg.Wait()
}

// ProcessURL parses the JSON data and stores document.
//
func (c *Crawler) ProcessURL(url string) (doc xkcdsearch.Document) {
	c.log.Debug().Str("URL", url).Msg("Processing URL")

	var id int

	if url != c.StartURL {
		id, err := c.documentIDFromURL(url)
		if err != nil {
			c.log.Fatal().Err(err).Msg("Error getting ID from URL")
		}

		if id == 404 {
			c.log.Info().Str("ID", "404").Msg("Skipping missing doc")
			return doc
		}

		if ok, err := c.existsDocument(strconv.Itoa(id)); ok {
			if err != nil {
				c.log.Error().Err(err).Int("ID", id).Msg("Error skipping existing doc")
			} else {
				c.log.Info().Int("ID", id).Msg("Skipping existing doc")
			}
			return doc
		}
	}

	res, err := c.loadURL(url)
	if err != nil {
		c.log.Error().Err(err).Int("ID", id).Msg("Error loading xkcd.com")
	}

	doc, err = c.processResponse(res)
	if err != nil {
		c.log.Error().Err(err).Str("ID", doc.ID).Msg("Error processing response")
	} else {
		err = c.storeDocument(&doc)
		if err != nil {
			c.log.Error().Err(err).Str("ID", doc.ID).Msg("Error storing doc")
		} else {
			c.log.Info().Str("ID", doc.ID).Str("title", doc.Title).Msg("Stored doc")
		}
	}

	time.Sleep(time.Duration(rand.Intn(100)+100) * time.Millisecond)

	return doc
}

// NextURL returns the next URL or an empty string.
//
func (c *Crawler) NextURL(url string) string {
	return c.nextURL(c, url)
}

func (c *Crawler) documentIDFromURL(url string) (id int, err error) {
	m := c.reURL.FindStringSubmatch(url)
	if len(m) > 0 && m[1] != "" {
		return strconv.Atoi(m[1])
	}
	return 0, fmt.Errorf("cannot get ID from URL [%s]", url)
}

func (c *Crawler) loadURL(url string) (*http.Response, error) {
	return http.Get(url)
}

func (c *Crawler) existsDocument(id string) (bool, error) {
	ok, err := c.store.Exists(id)
	if err != nil {
		return false, fmt.Errorf("store: %s", err)
	}
	return ok, nil
}

func (c *Crawler) storeDocument(doc *xkcdsearch.Document) error {
	return c.store.Create(doc)
}

func (c *Crawler) processResponse(res *http.Response) (xkcdsearch.Document, error) {
	defer res.Body.Close()

	var doc xkcdsearch.Document

	type jsonResponse struct {
		Num        int `json:"num"`
		Year       string
		Month      string
		Day        string
		Title      string `json:"title"`
		Transcript string
		Alt        string
		Link       string
		News       string
		Img        string
	}

	if res.StatusCode != 200 {
		return doc, fmt.Errorf("[%s]", res.Status)
	}

	var j jsonResponse
	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		return doc, err
	}

	jYear, err := strconv.ParseInt(j.Year, 10, 16)
	if err != nil {
		return doc, fmt.Errorf("strconv: %s", err)
	}
	jMonth, err := strconv.ParseInt(j.Month, 10, 8)
	if err != nil {
		return doc, fmt.Errorf("strconv: %s", err)
	}
	jDay, err := strconv.ParseInt(j.Day, 10, 8)
	if err != nil {
		return doc, fmt.Errorf("strconv: %s", err)
	}

	doc = xkcdsearch.Document{
		ID:         strconv.Itoa(j.Num),
		ImageURL:   j.Img,
		Published:  fmt.Sprintf("%04d-%02d-%02d", jYear, jMonth, jDay),
		Title:      j.Title,
		Alt:        j.Alt,
		Transcript: j.Transcript,
		Link:       j.Link,
		News:       j.News,
	}

	c.log.Debug().Interface("doc", doc).Msg("Downloaded")
	return doc, nil
}

func (c *Crawler) setupIndex() error {
	mapping := `{
    "mappings": {
      "_doc": {
        "properties": {
          "id":         { "type": "keyword" },
          "image_url":  { "type": "keyword" },
          "title":      { "type": "text", "analyzer": "english" },
          "alt":        { "type": "text", "analyzer": "english" },
          "transcript": { "type": "text", "analyzer": "english" },
          "published":  { "type": "date" },
          "link":       { "type": "keyword" },
          "news":       { "type": "text", "analyzer": "english" }
        }
      }
    }
		}`
	return c.store.CreateIndex(mapping)
}
