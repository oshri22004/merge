package github

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v50/github"
)

func WaitForWorkflows(prNumber int, repoName string) error {
	ctx := context.Background()
	client := NewClient(ctx)

	for {
		runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, os.Getenv("GITHUB_USERNAME"), repoName, &github.ListWorkflowRunsOptions{})
		if err != nil {
			return fmt.Errorf("failed to list workflow runs: %v", err)
		}

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
	return nil
}
