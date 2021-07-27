package version

import (
	"fmt"
)

var version = "unknown"
var commit = "unknown"

var Version = VersionContext{
	Version: version,
	Commit:  commit,
}

type VersionContext struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

func (vc *VersionContext) String() string {
	return fmt.Sprintf("%s <commit: %s>", vc.Version, vc.Commit)
}
