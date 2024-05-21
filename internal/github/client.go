package github

import (
	"context"
	"os"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func NewClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("MERGE_PAT")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
