package action

import (
	"github.com/Song-yu/glide/gom"
	"github.com/Song-yu/glide/msg"
)

// ImportGom imports a Gomfile.
func ImportGom(dest string) {
	base := "."
	config := EnsureConfig()
	if !gom.Has(base) {
		msg.Die("No gom data found.")
	}
	deps, err := gom.Parse(base)
	if err != nil {
		msg.Die("Failed to extract Gomfile: %s", err)
	}
	appendImports(deps, config)
	writeConfigToFileOrStdout(config, dest)
}
