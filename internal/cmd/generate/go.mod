module github.com/elastic/go-elasticsearch/v7/internal/cmd/generate

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../../../

require (
	github.com/alecthomas/chroma v0.6.3
	github.com/elastic/go-elasticsearch/v7 v7.0.0-20190407092644-3fb2a278216b
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	golang.org/x/tools v0.0.0-20200312194400-c312e98713c2
	gopkg.in/yaml.v2 v2.2.2
)
