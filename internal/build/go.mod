module github.com/elastic/go-elasticsearch/v8/internal/build

go 1.20

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/alecthomas/chroma v0.10.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/spf13/cobra v1.8.0
	golang.org/x/crypto v0.19.0
	golang.org/x/tools v0.18.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
)
