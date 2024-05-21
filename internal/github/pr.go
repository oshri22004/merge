package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v50/github"
)

const (
	baseBranch = "main"
	owner      = "DigitalIDF"
)

func OpenPullRequest(branch, repoName string) *github.PullRequest {
	ctx := context.Background()
	client := NewClient(ctx)

	newPR := &github.NewPullRequest{
		Title:               github.String("Automated PR"),
		Head:                github.String(branch),
		Base:                github.String(baseBranch),
		Body:                github.String("This is an automated PR."),
		MaintainerCanModify: github.Bool(true),
	}

	pr, _, err := client.PullRequests.Create(ctx, owner, repoName, newPR)
	if err != nil {
		log.Fatalf("Failed to create pull request: %v", err)
	}
	fmt.Printf("Opened PR: %s\n", pr.GetHTMLURL())

	return pr
}

func MergePullRequest(prNumber int, repoName string) {
	ctx := context.Background()
	client := NewClient(ctx)

	_, _, err := client.PullRequests.Merge(ctx, owner, repoName, prNumber, "Automated merge", &github.PullRequestOptions{})
	if err != nil {
		log.Fatalf("Failed to merge pull request: %v", err)
	}
	fmt.Println("Pull request merged successfully!")
}
