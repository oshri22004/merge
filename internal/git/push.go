package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Push(branch string) error {
	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push branch: %v", err)
	}
	return nil
}
