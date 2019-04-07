module github.com/elastic/go-elasticsearch/v6/internal/cmd/generate

go 1.11

replace github.com/elastic/go-elasticsearch/v6 => ../../../

require (
	github.com/elastic/go-elasticsearch/v6 6.x

	github.com/alecthomas/chroma v0.6.3
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/tools v0.0.0-20190330180304-aef51cc3777c
	gopkg.in/yaml.v2 v2.2.2
)
