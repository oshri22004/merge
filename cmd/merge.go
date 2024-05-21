package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

const (
	baseBranch = "main"
)

func main() {
	// Get the current branch
	branch, err := getCurrentBranch()
	if err != nil {
		log.Fatalf("Failed to get current branch: %v", err)
	}

	// Get the current repository name
	repoName, err := getCurrentRepositoryName()
	if err != nil {
		log.Fatalf("Failed to get current repository name: %v", err)
	}

	// Git push
	gitPush(branch)

	// Open PR
	pr := openPullRequest(branch, repoName)

	// Wait for workflows
	waitForWorkflows(pr.GetNumber(), repoName)

	// Merge PR
	mergePullRequest(pr.GetNumber(), repoName)
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	branch := strings.TrimSpace(out.String())
	return branch, nil
}

func getCurrentRepositoryName() (string, error) {
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

func gitPush(branch string) {
	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to push branch: %v", err)
	}
}

func openPullRequest(branch, repoName string) *github.PullRequest {
	ctx := context.Background()
	client := newGitHubClient(ctx)

	newPR := &github.NewPullRequest{
		Title:               github.String("Automated PR"),
		Head:                github.String(branch),
		Base:                github.String(baseBranch),
		Body:                github.String("This is an automated PR."),
		MaintainerCanModify: github.Bool(true),
	}

	pr, _, err := client.PullRequests.Create(ctx, "oshri22004", repoName, newPR)
	if err != nil {
		log.Fatalf("Failed to create pull request: %v", err)
	}
	fmt.Printf("Opened PR: %s\n", pr.GetHTMLURL())

	return pr
}

func waitForWorkflows(prNumber int, repoName string) {
	ctx := context.Background()
	client := newGitHubClient(ctx)

	for {
		runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, "oshri22004", repoName, &github.ListWorkflowRunsOptions{})
		if err != nil {
			log.Fatalf("Failed to list workflow runs: %v", err)
		}

		// Filter runs for the specific PR
		var prRuns []*github.WorkflowRun
		for _, run := range runs.WorkflowRuns {
			if run.PullRequests != nil {
				for _, pr := range run.PullRequests {
					if pr.GetNumber() == prNumber {
						prRuns = append(prRuns, run)
						break
					}
				}
			}
		}

		if len(prRuns) == 0 {
			break
		}

		fmt.Println("Waiting for workflows to complete...")
		time.Sleep(30 * time.Second)
	}
}

func mergePullRequest(prNumber int, repoName string) {
	ctx := context.Background()
	client := newGitHubClient(ctx)

	_, _, err := client.PullRequests.Merge(ctx, "oshri22004", repoName, prNumber, "Automated merge", &github.PullRequestOptions{})
	if err != nil {
		log.Fatalf("Failed to merge pull request: %v", err)
	}
	fmt.Println("Pull request merged successfully!")
}

func newGitHubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
