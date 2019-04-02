package dep

import (
	"context"
	"github.com/BurntSushi/toml"
	. "github.com/sanjid133/gopher-love/pkg"
	"path/filepath"
)

const (
	Manager  = "dep"
	FileName = "Gopkg.toml"
)

type Dep struct {
	ctx       context.Context
	directory string
}

var _ LoveBag = &Dep{}

func init() {
	RegistarManager(Manager, func(ctx context.Context) LoveBag { return New(ctx) })
}

func New(ctx context.Context) LoveBag {
	return &Dep{ctx: ctx}
}

type GopkgConfig struct {
	Constraint []constraint `toml:"constraint"`
	Prune      prune        `toml:"prune"`
}

type constraint struct {
	Name      string
	Source    string `toml:"-"`
	Branch    string `toml:"-"`
	Version   string `toml:"-"`
	Revisions string `toml:"-"`
}

type prune struct {
}

func (d *Dep) Initialize(directory string) LoveBag {
	d.directory = directory
	return d
}

func (d *Dep) File() string {
	return FileName
}

func (d *Dep) Read() ([]*Repository, error) {
	file := filepath.Join(d.directory, FileName)
	var config GopkgConfig
	_, err := toml.DecodeFile(file, &config)
	if err != nil {
		return nil, err
	}
	repos := make([]*Repository, 0)
	for _, c := range config.Constraint {
		repo := UrlToRepo(c.Name)
		repos = append(repos, repo)

	}
	return repos, nil
}
