package github

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	. "github.com/sanjid133/gopher-love/pkg/platform"
	"context"
	"fmt"
)

const Platform  = "githubs"

type Love struct {
	ctx context.Context
	client *github.Client

	orgName string

}

var _ MakeLove = &Love{}

func Initialize(token string) (*Love, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Love{
		ctx:ctx,
		client: github.NewClient(tc),
	}, nil

	// list all repositories for the authenticated user
	//repos, _, err := client.Repositories.List(ctx, "", nil)
}

func (l *Love) LoveOrganization(orgName string) error {
	repos, _, err := l.client.Repositories.List(l.ctx, orgName, &github.RepositoryListOptions{})
	if err != nil {
		return err
	}

	l.orgName = orgName
	for _, repo := range repos {
		fmt.Println(*repo.Name)
		l.sendLove(repo)
	}
	return nil

}

// If not starred a repository then star
func (l *Love) sendLove(repo *github.Repository)error  {
	starred, _, err := l.client.Activity.IsStarred(l.ctx, l.orgName, *repo.Name)
	if err != nil {
		return err
	}

	if !starred {
		if _, err = l.client.Activity.Star(l.ctx, l.orgName, *repo.Name); err != nil {
			return err
		}
	}
	return nil
}
