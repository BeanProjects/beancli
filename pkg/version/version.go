package version

import (
	"fmt"
	"runtime"
)


var (
	gitVersion   string = "v0.1.0"
	gitCommit    string = "$Format:%H$" 
	gitTreeState string = ""            
	buildDate 	 string = "1970-01-01T00:00:00Z" 
)


// Info contains versioning information.
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}



// Get returns an Info struct populated with version information.
// Assumes the variables gitMajor, gitMinor, etc., are defined elsewhere.
func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}


