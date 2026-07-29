module github.com/elastic/go-elasticsearch/v9/internal/build

go 1.25.0

toolchain go1.25.12

replace github.com/elastic/go-elasticsearch/v9 => ../../

require (
	github.com/alecthomas/chroma v0.10.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
	github.com/hashicorp/go-retryablehttp v0.7.8
	github.com/spf13/afero v1.15.0
	github.com/spf13/cobra v1.8.0
	golang.org/x/crypto v0.48.0
	golang.org/x/tools v0.42.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/term v0.40.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
