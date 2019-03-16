module github.com/elastic/go-elasticsearch/internal/cmd/generate

go 1.11

replace github.com/elastic/go-elasticsearch => ../../../

require (
	github.com/alecthomas/chroma v0.6.0
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/dlclark/regexp2 v1.1.6 // indirect
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/rs/zerolog v1.11.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/sys v0.0.0-20181128092732-4ed8d59d0b35 // indirect
	golang.org/x/tools v0.0.0-20181130052023-1c3d964395ce
	gopkg.in/yaml.v2 v2.2.2
)
