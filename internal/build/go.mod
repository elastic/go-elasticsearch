module github.com/elastic/go-elasticsearch/v8/internal/build

go 1.15

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/alecthomas/chroma v0.8.2
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220408190544-5352b0902921
	golang.org/x/tools v0.7.0
	gopkg.in/yaml.v2 v2.4.0
)
