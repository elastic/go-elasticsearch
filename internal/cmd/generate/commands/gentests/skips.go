package gentests

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

var skipTests map[string][]string

func init() {
	err := yaml.NewDecoder(strings.NewReader(skipTestsYAML)).Decode(&skipTests)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v", err))
	}
}

// TODO: Comments into descriptions for `Skip()`
//
var skipTestsYAML = `
---
# Cannot distinguish between missing value for refresh and an empty string
bulk/50_refresh.yaml:
  - refresh=empty string immediately makes changes are visible in search
create/60_refresh.yaml:
  - When refresh url parameter is an empty string that means "refresh immediately"
delete/50_refresh.yaml:
  - When refresh url parameter is an empty string that means "refresh immediately"
index/60_refresh.yaml:
  - When refresh url parameter is an empty string that means "refresh immediately"
update/60_refresh.yaml:
  - When refresh url parameter is an empty string that means "refresh immediately"

# Stash in value
cluster.reroute/11_explain.yaml:
nodes.info/30_settings.yaml:
nodes.stats/20_response_filtering.yaml:
nodes.stats/30_discovery.yaml:
 - Discovery stats
nodes.discovery/30_discovery.yaml:
 - Discovery stats

# Parsed response is YAML: value is map[interface {}]interface {}, not map[string]interface {}
cat.aliases/20_headers.yaml:
  - Simple alias with yaml body through Accept header

# Pattern breaks because of internal templates
cat.templates/10_basic.yaml:
  - Sort templates

# Not relevant
search/issue4895.yaml:
search/issue9606.yaml:

`
