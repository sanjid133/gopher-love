package pkg

import (
	"context"
	"os"
	"path/filepath"
	"github.com/sanjid133/gopher-love/pkg/system"
)

type Love interface {
	Initialize(config *system.SecretConfig) (Love, error)
	//Decode(url string) (*Repository, error)
	GetOrgRepos(org string)([]*Repository, error)
	SendLove(repo *Repository) error
}

type LoveBag interface {
	Initialize(directory string) LoveBag
	File()string
	Read()([]*Repository, error)
}

type Repository struct {
	Platform string
	Owner string
	Name string
}

func DetectManager(ctx context.Context, directory string) string  {
	//managers := GetAllRegistereredManager()
	for  m:= range manager {
		file := manager[m](ctx).File()
		if _, err := os.Stat(filepath.Join(directory, file)); err == nil {
			return m
		}
	}
	return ""
}