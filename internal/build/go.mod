module github.com/elastic/go-elasticsearch/v8/internal/build

go 1.20

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/alecthomas/chroma v0.8.2
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/spf13/cobra v1.1.3
	golang.org/x/crypto v0.0.0-20220408190544-5352b0902921
	golang.org/x/tools v0.7.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/dlclark/regexp2 v1.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
)
