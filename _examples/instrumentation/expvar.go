// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"expvar"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	// Import the "expvar" and "pprof" package >>>>>>>>>>
	"net/http"
	_ "net/http/pprof"
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
)

var (
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
)

func init() {
	runtime.SetMutexProfileFraction(10)
}

func main() {
	log.SetFlags(0)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	aborted := make(chan os.Signal)
	signal.Notify(aborted, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-aborted

		log.Println("\nDone!\n")
		os.Exit(0)
	}()

	log.Println(strings.Repeat("─", tWidth))
	log.Println("Open <http://localhost:6060/debug/vars> to see all exported variables.")
	log.Println(strings.Repeat("─", tWidth))

	// Start the debug server >>>>>>>>>>>>>>>>>>>>>>>>>>>
	go func() { log.Fatalln(http.ListenAndServe("localhost:6060", nil)) }()
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	for i := 1; i <= 2; i++ {
		go func(i int) {
			log.Printf("==> Starting server on <localhost:1000%d>", i)
			if err := http.ListenAndServe(
				fmt.Sprintf("localhost:1000%d", i),
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "OK\n") }),
			); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Unable to start server: %s", err)
			}
		}(i)
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{
				"http://localhost:10001",
				"http://localhost:10002",
				"http://localhost:10003",
			},

			Logger:            &estransport.ColorLogger{Output: os.Stdout},
			DisableRetry:      true,
			EnableDebugLogger: true,

			// Enable metric collection >>>>>>>>>>>>>>>>>>>>>>>>>
			EnableMetrics: true,
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
		})
	if err != nil {
		log.Fatal("ERROR: %s", err)
	}

	// Publish client metrics to expvar >>>>>>>>>>>>>>>>>
	expvar.Publish("go-elasticsearch", expvar.Func(func() interface{} { m, _ := es.Metrics(); return m }))
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	for {
		select {
		case t := <-ticker.C:

			go func() {
				if res, _ := es.Info(); res != nil {
					res.Body.Close()
				}
			}()

			go func() {
				if res, _ := es.Cluster.Health(); res != nil {
					res.Body.Close()
				}
			}()

			if t.Second()%5 == 0 {
				m, err := es.Metrics()
				if err != nil {
					log.Printf("\x1b[31;1mUnable to get metrics: %s\x1b[0m", err)
					continue
				}
				log.Println("███", fmt.Sprintf("\x1b[1m%s\x1b[0m", "Metrics"), strings.Repeat("█", tWidth-12))
				log.Printf(
					""+
						"    \x1b[2mRequests:   \x1b[0m %d\n"+
						"    \x1b[2mFailures:   \x1b[0m %d\n"+
						"    \x1b[2mConnections:\x1b[0m %s",
					m.Requests, m.Failures, m.Connections)
				log.Println(strings.Repeat("─", tWidth))
			}
		}
	}
}
