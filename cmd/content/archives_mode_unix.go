//go:build !windows

package content

import "io/fs"

// archiveFileModes returns a function resolving the permission bits of the
// archived files in folder. On Unix-like systems, the local file mode is the
// source of truth.
func archiveFileModes(folder string) func(rel string, info fs.FileInfo) int64 {
	return func(_ string, info fs.FileInfo) int64 {
		return int64(info.Mode().Perm())
	}
}
