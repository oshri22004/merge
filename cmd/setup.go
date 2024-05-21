package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Initial setup for GitHub credentials",
	Run: func(cmd *cobra.Command, args []string) {
		ensureSetup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

func ensureSetup() {
	if os.Getenv("GITHUB_USERNAME") == "" || os.Getenv("MERGE_PAT") == "" {
		fmt.Println("Initial setup required.")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your GitHub username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter your GitHub personal access token: ")
		token, _ := reader.ReadString('\n')
		token = strings.TrimSpace(token)

		// Save to ~/.zshrc
		zshrcPath := filepath.Join(os.Getenv("HOME"), ".zshrc")
		zshrcFile, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Failed to open ~/.zshrc: %v", err)
		}
		defer zshrcFile.Close()

		_, err = zshrcFile.WriteString(fmt.Sprintf("\nexport GITHUB_USERNAME=%s\nexport MERGE_PAT=%s\n", username, token))
		if err != nil {
			log.Fatalf("Failed to write to ~/.zshrc: %v", err)
		}

		// Load the new environment variables
		cmd := exec.Command("zsh", "-c", "source ~/.zshrc")
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Failed to source ~/.zshrc: %v", err)
		}

		fmt.Println("Setup complete. Please restart your terminal session.")
		os.Exit(0)
	}
}
