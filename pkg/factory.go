package pkg

import (
	"context"
	"github.com/sanjid133/gopher-love/pkg/system"
	"os"
	"path/filepath"
)

type Love interface {
	Initialize(config *system.SecretConfig) (Love, error)
	//Decode(url string) (*Repository, error)
	GetOrgRepos(org string) ([]*Repository, error)
	IsLoved(repo *Repository) (bool, error)
	SendLove(repo *Repository) error
}

type LoveBag interface {
	Initialize(directory string) LoveBag
	File() string
	Read() ([]*Repository, error)
}

type Repository struct {
	Platform string
	Owner    string
	Name     string

	Url string
}

func DetectManager(ctx context.Context, directory string) string {
	//managers := GetAllRegistereredManager()
	for m := range manager {
		file := manager[m](ctx).File()
		if _, err := os.Stat(filepath.Join(directory, file)); err == nil {
			return m
		}
	}
	return ""
}

var reaction = []string{
	":kissing_heart:",
	":kissing_closed_eyes:",
	":kissing_smiling_eyes:",
	":yellow_heart:",
	":blue_heart:",
	":purple_heart:",
	":heart:",
	":green_heart:",
	":heartpulse:",
	":heartbeat:",
	":two_hearts:",
	":revolving_hearts:",
	":sparkling_heart:",
	":gift_heart:",
}
