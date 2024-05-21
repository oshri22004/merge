package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentRepositoryName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get repository root: %w", err)
	}
	rootPath := strings.TrimSpace(out.String())
	parts := strings.Split(rootPath, string(os.PathSeparator))
	repoName := parts[len(parts)-1]
	return repoName, nil
}
