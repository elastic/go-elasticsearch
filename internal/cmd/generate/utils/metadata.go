package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	reEsVersion = regexp.MustCompile(`elasticsearch\s+=\s+(\d+\.\d+\.\d+)`)
)

// EsVersion returns the Elasticsearch from Java property file, or an error.
//
func EsVersion(fpath string) (string, error) {
	basePath, err := basePathFromFilepath(fpath)
	if err != nil {
		return "", fmt.Errorf("EsVersion: %s", err)
	}

	c, err := ioutil.ReadFile(filepath.Join(basePath, "buildSrc", "version.properties"))
	if err != nil {
		return "", fmt.Errorf("EsVersion: %s", err)
	}

	m := reEsVersion.FindSubmatch(c)
	if len(m) < 2 {
		return "", nil
	}
	return string(m[1]), nil
}

// GitCommit returns the Git commit for fpath, or an error.
//
func GitCommit(fpath string) (string, error) {
	if gitCommitEnv := os.Getenv("ELASTICSEARCH_BUILD_HASH"); gitCommitEnv != "" {
		return gitCommitEnv, nil
	}

	basePath, err := basePathFromFilepath(fpath)
	if err != nil {
		return "", fmt.Errorf("GitCommit: %s", err)
	}

	args := strings.Split("git --git-dir="+basePath+".git rev-parse --short HEAD", " ")
	cmd := exec.Command(args[0:1][0], args[1:]...)
	// fmt.Printf("> %s\n", strings.Join(cmd.Args, " "))

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("GitCommit: %s", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// GitTag returns the Git tag for fpath if available, or an error.
//
func GitTag(fpath string) (string, error) {
	basePath, err := basePathFromFilepath(fpath)
	if err != nil {
		return "", fmt.Errorf("GitCommit: %s", err)
	}

	commit, err := GitCommit(fpath)
	if err != nil {
		return "", fmt.Errorf("GitTag: %s", err)
	}

	args := strings.Split("git --git-dir="+basePath+".git tag --points-at "+commit, " ")
	cmd := exec.Command(args[0:1][0], args[1:]...)
	// fmt.Printf("> %s\n", strings.Join(cmd.Args, " "))

	out, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(out)), nil
}

func basePathFromFilepath(fpath string) (string, error) {
	var bpath strings.Builder

	fpath, err := filepath.Abs(fpath)
	if err != nil {
		return "", err
	}

	for _, p := range strings.Split(fpath, string(filepath.Separator)) {
		if p == "rest-api-spec" {
			break
		}
		bpath.WriteString(p)
		bpath.WriteRune(filepath.Separator)
	}

	return bpath.String(), nil
}
