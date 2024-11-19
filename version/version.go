package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type BuildInfo struct {
	GitTag       string `json:"git_tag,omitempty"`
	GitCommit    string `json:"git_commit,omitempty"`
	GitBranch    string `json:"git_branch,omitempty"`
	GitTreeState string `json:"git_tree_state,omitempty"`
	BuildDate    string `json:"build_date,omitempty"`
	BuiltBy      string `json:"built_by,omitempty"`
	GoVersion    string `json:"go_version,omitempty"`
	Compiler     string `json:"compiler,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Version      string `json:"version,omitempty"`
}

// go build -ldflags "
// -X github.com/origadmin/toolkits/version.gitTag=${TAG}
// -X github.com/origadmin/toolkits/version.gitCommit=${COMMIT}
// -X github.com/origadmin/toolkits/version.gitBranch=${BRANCH}
// -X github.com/origadmin/toolkits/version.buildDate=${DATE}
// -X github.com/origadmin/toolkits/version.builtBy=${BUILT_BY}
// -X github.com/origadmin/toolkits/version.version=${VERSION}
//
// ModulePath = github.com/origadmin/toolkits/version
//
// # gitHash The current commit id is the same as the gitCommit result
// gitHash = $(shell git rev-parse HEAD)
// gitBranch = $(shell git rev-parse --abbrev-ref HEAD)
// gitTag = $(shell \
// if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ]; \
// then \
// git describe --tags --abbrev=0; \
// else \
// git log --pretty=format:'%h' -n 1; \
// fi)
// # same as gitHash previously
// gitCommit = $(shell git log --pretty=format:'%H' -n 1)
//
// gitTreeState = $(shell if git status | grep -q 'clean'; then echo clean; else echo dirty; fi)
//
// buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
// # buildDate = $(shell TZ=Asia/Shanghai date +%F\ %T%z | tr 'T' ' ')
// LDFLAGS += -X "${ModulePath}.gitTag=${gitTag}"
// LDFLAGS += -X "${ModulePath}.buildDate=${buildDate}"
// LDFLAGS += -X "${ModulePath}.gitCommit=${gitCommit}"
// LDFLAGS += -X "${ModulePath}.gitTreeState=${gitTreeState}"
// LDFLAGS += -X "${ModulePath}.gitBranch=${gitBranch}"
// LDFLAGS += -X "${ModulePath}.version=${VERSION}"
var (
	gitTag       string = ""
	gitCommit    string = ""
	gitBranch    string = ""
	gitTreeState string = ""
	buildDate    string = time.Now().Format(time.RFC3339)
	builtBy      string = "OrigAdmin"
	version      string = "v1.0.0"
)

var (
	info BuildInfo
	once sync.Once
	text []byte
)

// initialize initializes the build information and marshals it to JSON.
func initialize() {
	info = BuildInfo{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitBranch:    gitBranch,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		BuiltBy:      builtBy,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Version:      version,
	}
	var err error
	text, err = json.MarshalIndent(info, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal build info during initialization: %v\n", err)
	}
}

// ReadBuildInfo returns an instance of BuildInfo.
// This function ensures that the build information is initialized only once using sync.Once,
// which improves performance and ensures thread safety.
//
// Returns:
//   - BuildInfo: A struct containing detailed build information.
func ReadBuildInfo() BuildInfo {
	once.Do(initialize)
	return info
}

// PrintBuildInfo prints the build information in JSON format.
// This function uses the sync.Once mechanism to ensure that the build information is initialized only once,
// thus avoiding the problem of concurrent initialization and ensuring thread safety.
func PrintBuildInfo() {
	once.Do(initialize)
	if text == nil {
		fmt.Println("marshaled build info is not available")
		return
	}
	fmt.Println(string(text))
}
