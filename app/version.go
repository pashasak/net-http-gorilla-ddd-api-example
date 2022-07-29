package app

// BuildInfo is a variable for storing build information.
type BuildInfo struct {
	BuildTime    string `json:"buildTime"`
	BuildBranch  string `json:"buildBranch"`
	BuildCommit  string `json:"buildCommit"`
	BuildSummary string `json:"buildSummary"`
}

var (
	// BuildTime is n output of "date +"%Y.%m.%d-%T.%Z""
	BuildTime = ""
	// BuildBranch is an output of "git symbolic-ref  --short HEAD"
	BuildBranch = ""
	// BuildCommit is an output of "git rev-parse HEAD"
	BuildCommit = ""
	// BuildSummary is an output of "git describe --tags --dirty --always"
	BuildSummary = ""
)
