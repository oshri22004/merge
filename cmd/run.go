/*
Copyright © 2024 NAME HERE oshripar@gmail.com
*/

package cmd

import (
	"log"
	"time"

	"github.com/oshri22004/merge/internal/git"
	"github.com/oshri22004/merge/internal/github"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <commit message>",
	Short: "Run the GitHub merge automation",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commitMessage := args[0]
		ensureSetup()

		branch, err := git.GetCurrentBranch()
		if err != nil {
			log.Fatalf("Failed to get current branch: %v", err)
		}

		repoName, err := git.GetCurrentRepositoryName()
		if err != nil {
			log.Fatalf("Failed to get current repository name: %v", err)
		}

		git.Add()
		git.Commit(commitMessage)
		git.Push(branch)
		pr := github.OpenPullRequest(branch, repoName)
		time.Sleep(15 * time.Second)
		github.MergePullRequest(pr.GetNumber(), repoName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
