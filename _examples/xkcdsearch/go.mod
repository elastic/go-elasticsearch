module github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 master
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/rs/zerolog v1.11.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/sys v0.0.0-20181128092732-4ed8d59d0b35 // indirect
)
