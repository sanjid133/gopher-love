package glide

import (
	"context"
	. "github.com/sanjid133/gopher-love/pkg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	Manager  = "glide"
	FileName = "glide.yaml"
)

type Glide struct {
	ctx       context.Context
	directory string
}

var _ LoveBag = &Glide{}

func init() {
	RegistarManager(Manager, func(ctx context.Context) LoveBag { return New(ctx) })
}

func New(ctx context.Context) LoveBag {
	return &Glide{ctx: ctx}
}

type GlideFile struct {
	Package string `yaml:"package"`
	Import  []struct {
		Package string `yaml:"package"`
		Repo    string `yaml:"repo,omitempty"`
	} `yaml:"import"`
}

func (d *Glide) Initialize(directory string) LoveBag {
	d.directory = directory
	return d
}

func (d *Glide) File() string {
	return FileName
}

func (d *Glide) Read() ([]*Repository, error) {
	file := filepath.Join(d.directory, FileName)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &GlideFile{}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	repos := make([]*Repository, 0)
	for _, c := range config.Import {
		repo := UrlToRepo(c.Package)
		repos = append(repos, repo)

	}
	return repos, nil
}
