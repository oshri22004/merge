package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to commit -m %s: %v", message, err)
	}
	return nil
}
