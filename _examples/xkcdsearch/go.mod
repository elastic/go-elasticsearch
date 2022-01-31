module github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch

go 1.17

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/rs/zerolog v1.11.0
	github.com/spf13/cobra v0.0.3
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.0.0-20211216131617-bbee439d559c // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20181128092732-4ed8d59d0b35 // indirect
)
