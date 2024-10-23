package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"
)

var (
	gitTag       string = ""
	gitCommit    string = ""
	gitBranch    string = ""
	gitTreeState string = ""
	buildDate    string = ""
	version      string = ""
)

type BuildInfo struct {
	GitTag       string `json:"git_tag"`
	GitCommit    string `json:"git_commit"`
	GitBranch    string `json:"git_branch"`
	GitTreeState string `json:"git_tree_state"`
	BuildDate    string `json:"build_date"`
	GoVersion    string `json:"go_version"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
	Version      string `json:"version"`
}

var info func() BuildInfo

func init() {
	info = sync.OnceValue(func() BuildInfo {
		return BuildInfo{
			GitTag:       gitTag,
			GitCommit:    gitCommit,
			GitBranch:    gitBranch,
			GitTreeState: gitTreeState,
			BuildDate:    buildDate,
			GoVersion:    runtime.Version(),
			Compiler:     runtime.Compiler,
			Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			Version:      version,
		}
	})
}

func Get() BuildInfo {
	return info()
}

func Show() {
	v := Get()
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(marshalled))
}
