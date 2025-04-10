module github.com/elastic/go-elasticsearch/v9/internal/build

go 1.23.0

toolchain go1.24.2

replace github.com/elastic/go-elasticsearch/v9 => ../../

require (
	github.com/alecthomas/chroma v0.10.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.0
	golang.org/x/crypto v0.19.0
	golang.org/x/tools v0.32.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/term v0.17.0 // indirect
)
