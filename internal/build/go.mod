module github.com/elastic/go-elasticsearch/v7/internal/build

go 1.15

replace github.com/elastic/go-elasticsearch/v7 => ../../

require (
	github.com/alecthomas/chroma v0.8.2
	github.com/elastic/go-elasticsearch/v7 v7.5.1-0.20210608143310-9747f0e42f35
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	golang.org/x/tools v0.1.0
	gopkg.in/yaml.v2 v2.4.0
)
