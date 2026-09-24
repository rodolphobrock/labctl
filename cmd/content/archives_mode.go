package content

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// gitIndexModes returns the file modes (e.g. "100644", "100755") recorded in
// the git index for the files under folder, keyed by slash-separated paths
// relative to folder.
func gitIndexModes(folder string) (map[string]string, error) {
	out, err := exec.Command("git", "-C", folder, "ls-files", "--stage", "-z").Output()
	if err != nil {
		return nil, fmt.Errorf("git ls-files: %w", err)
	}
	return parseGitIndexModes(out)
}

// parseGitIndexModes parses `git ls-files --stage -z` output, which is a
// sequence of NUL-terminated "<mode> <object> <stage>\t<path>" entries.
func parseGitIndexModes(out []byte) (map[string]string, error) {
	modes := map[string]string{}
	for _, entry := range bytes.Split(out, []byte{0}) {
		if len(entry) == 0 {
			continue
		}

		meta, path, ok := strings.Cut(string(entry), "\t")
		if !ok {
			return nil, fmt.Errorf("malformed git ls-files entry: %q", entry)
		}
		mode, _, _ := strings.Cut(meta, " ")
		modes[path] = mode
	}
	return modes, nil
}
