package build

import "fmt"

// go build -ldflags "-X github.com/xbmlz/gin-svelte-template/cmd.Version=x.y.z"
var (
	// Version is the version of the project
	Version = "0.0.0"
	// Revision is the git short commit revision number
	// If built without a Git repository, this field will be empty.
	Revision = ""
	// Time is the build time of the project
	Time = ""
)

func String() string {
	return fmt.Sprintf("%s\nrevision: %s\ntime: %s\n", Version, Revision, Time)
}
