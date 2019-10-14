// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/go-elasticsearch/v8"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

const count = 100

var (
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))

	faint   = color.New(color.Faint)
	bold    = color.New(color.Bold)
	boldRed = color.New(color.FgRed).Add(color.Bold)

	tagMethod, _ = tag.NewKey("method")
)

func init() {
	if tWidth < 0 {
		tWidth = 0
	}
}

// ConsoleExporter prints stats and traces to STDOUT.
//
type ConsoleExporter struct{}

// ExportView prints the stats.
//
func (e *ConsoleExporter) ExportView(vd *view.Data) {
	fmt.Println(strings.Repeat("─", tWidth))
	for _, row := range vd.Rows {
		faint.Print("█ ")
		fmt.Printf("%-17s ", strings.TrimPrefix(vd.View.Name, "opencensus.io/http/client/"))

		switch v := row.Data.(type) {
		case *view.DistributionData:
			bold.Printf("min=%-6.1f max=%-6.1f mean=%-6.1f", v.Min, v.Max, v.Mean)
		case *view.CountData:
			bold.Printf("count=%-3v", v.Value)
		case *view.SumData:
			bold.Printf("sum=%-3v", v.Value)
		case *view.LastValueData:
			bold.Printf("last=%-3v", v.Value)
		}
		faint.Print(" │ ")

		for _, tag := range row.Tags {
			faint.Printf("%-25s ", fmt.Sprintf("%v=%v", tag.Key.Name(), tag.Value))
		}
		fmt.Println()
	}
}

// ExportSpan prints the traces.
//
func (e *ConsoleExporter) ExportSpan(sd *trace.SpanData) {
	var c *color.Color
	if sd.Status.Code > 0 {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgGreen)
	}

	fmt.Println(strings.Repeat("─", tWidth))

	fmt.Printf(
		"░ %s %s %s\n",
		c.Sprint(sd.Status.Message),
		bold.Sprint(sd.Name),
		sd.EndTime.Sub(sd.StartTime).Round(time.Millisecond))

	faint.Printf("░ %x > %x\n", sd.SpanContext.TraceID[:], sd.SpanContext.SpanID[:])

	if len(sd.Attributes) > 0 {
		faint.Print("░ ")
		var keys []string
		for k, _ := range sd.Attributes {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			faint.Printf("%s=%v", k, sd.Attributes[k])
			if i < len(keys)-1 {
				faint.Printf(" │ ")
			}
		}
	}
	fmt.Println()
}

func main() {
	log.SetFlags(0)
	start := time.Now()

	// Create new elasticsearch client ...
	//
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// ... using the "ochttp" wrapper for instrumentation
			Transport: &ochttp.Transport{},
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
		})
	if err != nil {
		log.Fatal("ERROR: %s", err)
	}

	// Set up the "done" channel
	//
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	// Set up tickers
	//
	tickers := struct {
		Info   *time.Ticker
		Index  *time.Ticker
		Health *time.Ticker
		Search *time.Ticker
	}{
		Info:   time.NewTicker(time.Second),
		Index:  time.NewTicker(500 * time.Millisecond),
		Health: time.NewTicker(5 * time.Second),
		Search: time.NewTicker(10 * time.Second),
	}
	defer tickers.Info.Stop()
	defer tickers.Index.Stop()
	defer tickers.Health.Stop()
	defer tickers.Search.Stop()

	// Register views bundled with the "ochttp" plugin
	//
	if err := view.Register(
		ochttp.ClientRoundtripLatencyDistribution,
		ochttp.ClientCompletedCount,
	); err != nil {
		log.Fatal("ERROR: %s", err)
	}

	// Report views to STDOUT once in a while
	//
	view.SetReportingPeriod(5 * time.Second)
	view.RegisterExporter(&ConsoleExporter{})

	// Report a portion of traces to STDOUT
	//
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(0.5)})
	trace.RegisterExporter(&ConsoleExporter{})

	// Initialize the context
	//
	ctx, _ := tag.New(context.Background(), tag.Upsert(tagMethod, "main"))

	// Perform API calls
	//
	for {
		select {
		case <-done:
			fmt.Print("\n")
			fmt.Println(strings.Repeat("━", tWidth))
			faint.Printf("Finished in %s\n\n", time.Now().Sub(start).Truncate(time.Second))
			return

		// -> Info
		//
		case <-tickers.Info.C:
			res, err := es.Info(es.Info.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				res.Body.Close()
			}

		// -> Index
		//
		case t := <-tickers.Index.C:
			// Artificially fail some requests...
			var body io.Reader
			if t.Second()%4 == 0 {
				body = strings.NewReader(``)
			} else {
				body = strings.NewReader(`{"timestamp":"` + t.Format(time.RFC3339) + `"}`)
			}

			res, err := es.Index("test", body, es.Index.WithContext(ctx))
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				res.Body.Close()
			}

		// -> Health
		//
		case <-tickers.Health.C:
			res, err := es.Cluster.Health(
				es.Cluster.Health.WithLevel("indices"),
				es.Cluster.Health.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				res.Body.Close()
			}

		// -> Search
		//
		case <-tickers.Search.C:
			res, err := es.Search(
				es.Search.WithIndex("test"),
				es.Search.WithSort("timestamp:desc"),
				es.Search.WithSize(1),
				es.Search.WithContext(ctx),
			)
			if err != nil {
				boldRed.Printf("Error getting response: %s\n", err)
			} else {
				res.Body.Close()
			}
		}
	}
}
