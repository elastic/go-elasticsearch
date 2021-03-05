module github.com/elastic/go-elasticsearch/v8/internal/cmd/generate

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../../../

require (
	github.com/alecthomas/chroma v0.6.3
	github.com/elastic/go-elasticsearch/v8 master
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.1.0
	gopkg.in/yaml.v2 v2.2.2
)
