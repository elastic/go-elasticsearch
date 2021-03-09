package estransport

import (
	"regexp"
	"runtime"
	"strings"
)

// HeaderClientMeta Key for the HTTP Header related to telemetry data sent with
// each request to Elasticsearch.
const HeaderClientMeta = "x-elastic-client-meta"

var metaReVersion = regexp.MustCompile("([0-9.]+)(.*)")

func initMetaHeader() string {
	var b strings.Builder
	var strippedGoVersion string
	var strippedEsVersion string

	strippedEsVersion = buildStrippedVersion(Version)
	strippedGoVersion = buildStrippedVersion(runtime.Version())

	var duos = [][]string{
		{
			"es",
			strippedEsVersion,
		},
		{
			"go",
			strippedGoVersion,
		},
		{
			"t",
			strippedEsVersion,
		},
		{
			"hc",
			strippedGoVersion,
		},
	}

	var arr []string
	for _, duo := range duos {
		arr = append(arr, strings.Join(duo, "="))
	}
	b.WriteString(strings.Join(arr, ","))

	return b.String()
}

func buildStrippedVersion(version string) string {
	v := metaReVersion.FindStringSubmatch(version)

	if len(v) == 3 && !strings.Contains(version, "devel") {
		switch {
		case v[2] != "":
			return v[1] + "p"
		default:
			return v[1]
		}
	}

	return "0.0p"
}
