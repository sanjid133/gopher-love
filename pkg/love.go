package pkg

import (
	"fmt"
	"strings"
	"context"
	"github.com/sanjid133/gopher-love/pkg/system"
	"github.com/sanjid133/gopher-love/util"
)

func LoveOrganization(url string) error {
	parts := strings.Split(url, "/")
	if len(parts) != 2 {
		return fmt.Errorf("Org name %v is not valid. Correct format is github.com/<org>", url)
	}
	ctx := context.Background()

	org := util.GetPlatform(parts[0])
	platform, err := GetPlatform(org, ctx)
	if err != nil {
		return err
	}
	love, err := platform.Initialize(system.Config)
	if err != nil {
		return err
	}
	
	repos, err := love.GetOrgRepos(parts[1])
	if err != nil {
		return err
	}
	
	return LoveRepos(love, repos)
}

func LoveDependency(directory string)error  {
	ctx := context.Background()
	
	dependencyType := DetectManager(ctx, directory)
	
	manager, err := GetManager(dependencyType, ctx)
	if err != nil {
		return err
	}
	repos, err := manager.Read()
	if err != nil {
		return err
	}

	orgRepos := SortOrganization(repos)
	for org, repos := range orgRepos {
		platform, err := GetPlatform(org, ctx)
		if err != nil {
			return err
		}
		love, err := platform.Initialize(system.Config)
		if err != nil {
			return err
		}
		if err = LoveRepos(love, repos); err != nil {
			return err
		}
	}
	return nil
}

func LoveRepos(love Love, repos []*Repository) error  {
	for _, repo := range repos {
		if err := love.SendLove(repo); err != nil{
			return err
		}
	}
	return nil
}

func SortOrganization(repos []*Repository) map[string][]*Repository {
	output := make(map[string][]*Repository)
	for _, repo := range repos {
		output[repo.Platform] = append(output[repo.Platform], repo)
	}
	return output
}