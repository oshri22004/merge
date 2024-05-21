package git

import (
	"log"
	"os"
	"os/exec"
)

func Add() {
	cmd := exec.Command("git", "add", "-A")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to git add -A(all): %v", err)
	}
}
