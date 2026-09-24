//go:build windows

package content

import (
	"io/fs"
	"log/slog"
	"sync"
)

// warnedNoGitIndex keeps watch mode from repeating the same warning on every
// archive rebuild.
var warnedNoGitIndex sync.Map

// archiveFileModes returns a function resolving the permission bits of the
// archived files in folder. Windows doesn't track Unix permissions, so the
// executable bit comes from the git index instead (set it with
// `git update-index --chmod=+x <file>`). Files not tracked by git get 0644.
func archiveFileModes(folder string) func(rel string, info fs.FileInfo) int64 {
	modes, err := gitIndexModes(folder)
	if err != nil {
		if _, warned := warnedNoGitIndex.LoadOrStore(folder, true); !warned {
			slog.Warn("Couldn't read file modes from the git index; archiving all files as non-executable (0644)",
				"folder", folder, "error", err)
		}
	}

	return func(rel string, _ fs.FileInfo) int64 {
		if modes[rel] == "100755" {
			return 0755
		}
		return 0644
	}
}
