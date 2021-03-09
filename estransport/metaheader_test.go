package estransport

import (
	"regexp"
	"runtime"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/internal/version"
)

var (
	metaHeaderReValidation = regexp.MustCompile(`^[a-z]{1,}=[a-z0-9\.\-]{1,}(?:,[a-z]{1,}=[a-z0-9\.\-]+)*$`)
)

func Test_buildStrippedVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Standard Go version",
			args: args{version: "go1.16"},
			want: "1.16",
		},
		{
			name: "Rc Go version",
			args: args{
				version: "go1.16rc1",
			},
			want: "1.16p",
		},
		{
			name: "Beta Go version (go1.16beta1 example)",
			args: args{
				version: "devel +2ff33f5e44 Thu Dec 17 16:03:19 2020 +0000",
			},
			want: "0.0p",
		},
		{
			name: "Random mostly good Go version",
			args: args{
				version: "1.16",
			},
			want: "1.16",
		},
		{
			name: "Client package version",
			args: args{
				version: "8.0.0",
			},
			want: "8.0.0",
		},
		{
			name: "Client pre release version",
			args: args{
				version: "8.0.0-SNAPSHOT",
			},
			want: "8.0.0p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStrippedVersion(tt.args.version); got != tt.want {
				t.Errorf("buildStrippedVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initMetaHeader(t *testing.T) {
	esVersion := buildStrippedVersion(version.Client)
	goVersion := buildStrippedVersion(runtime.Version())

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Meta header generation",
			want: "es=" + esVersion + ",go=" + goVersion + ",t=" + esVersion + ",hc=" + goVersion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if metaHeader != tt.want {
				t.Errorf("initMetaHeader() = %v, want %v", metaHeader, tt.want)
			}
			if valid := metaHeaderReValidation.Match([]byte(metaHeader)); valid != true {
				t.Errorf("Metaheder doesn't validate regexp format validation, got : %s", metaHeader)
			}
		})
	}
}
