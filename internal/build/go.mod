module github.com/elastic/go-elasticsearch/v9/internal/build

go 1.24.0

toolchain go1.24.2

replace github.com/elastic/go-elasticsearch/v9 => ../../

require (
	github.com/alecthomas/chroma v0.10.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.0
	golang.org/x/crypto v0.37.0
	golang.org/x/tools v0.39.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.30.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/telemetry v0.0.0-20251111182119-bc8e575c7b54 // indirect
	golang.org/x/term v0.31.0 // indirect
)
