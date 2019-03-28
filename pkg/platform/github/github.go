package github

import (
	"context"
	"github.com/google/go-github/github"
	. "github.com/sanjid133/gopher-love/pkg"
	"github.com/sanjid133/gopher-love/pkg/system"
	"golang.org/x/oauth2"
)

const Platform = "github"

type Github struct {
	ctx    context.Context
	client *github.Client
}

var _ Love = &Github{}

func init() {
	RegistarPlatform(Platform, func(ctx context.Context) (Love, error) { return New(ctx), nil })
}

func New(ctx context.Context) Love {
	return &Github{ctx: ctx}

}
func (g *Github) Initialize(config *system.SecretConfig) (Love, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Github.ApiToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	g.client = github.NewClient(tc)
	return g, nil
}

func (g *Github) GetOrgRepos(org string) ([]*Repository, error) {
	repos, _, err := g.client.Repositories.List(g.ctx, org, &github.RepositoryListOptions{})
	if err != nil {
		return nil, err
	}
	retRepos := make([]*Repository, 0)
	for _, repo := range repos {
		retRepos = append(retRepos, &Repository{
			Platform: Platform,
			Owner:    org,
			Name:     *repo.Name,
		})
	}
	return retRepos, err
}

func (g *Github) SendLove(repo *Repository) error {
	starred, _, err := g.client.Activity.IsStarred(g.ctx, repo.Owner, repo.Name)
	if err != nil {
		return err
	}

	if !starred {
		if _, err = g.client.Activity.Star(g.ctx, repo.Owner, repo.Name); err != nil {
			return err
		}
	}
	return nil
}
