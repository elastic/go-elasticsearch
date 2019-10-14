// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package commands

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch"
)

var (
	listenHost = "localhost"
	listenPort = "8000"
	listenAddr string
)

func init() {
	if envListenHost := os.Getenv("HTTP_LISTEN_HOST"); envListenHost != "" {
		listenHost = envListenHost
	}
	if envListenPort := os.Getenv("HTTP_LISTEN_PORT"); envListenPort != "" {
		listenPort = envListenPort
	}
	listenAddr = listenHost + ":" + listenPort

	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&listenAddr, "listen", "l", listenAddr, "Listening address")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Launch a web interface",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating Elasticsearch client")
		}

		config := xkcdsearch.StoreConfig{Client: es, IndexName: IndexName}
		store, err := xkcdsearch.NewStore(config)
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating store")
		}
		http.Handle("/search.json", store)
		http.Handle("/", http.FileServer(http.Dir("./web")))

		log.Info().Msgf("Starting server at <%s>...", listenAddr)
		err = http.ListenAndServe(listenAddr, nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Error launching server")
		}
	},
}
