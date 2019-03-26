package dep

import (
	"C"
	"github.com/BurntSushi/toml"
	"github.com/sanjid133/gopher-love/util"
	"fmt"
	"path/filepath"
)

type GopkgConfig struct {
	Constraint []constraint `toml:"constraint"`
	Prune prune `toml:"prune"`
}

type constraint struct {
	Name string
	Source string `toml:"-"`
	Branch string `toml:"-"`
	Version string `toml:"-"`
	Revisions string `toml:"-"`
}

type prune struct {
}

func Parse(file string) error {
	var config GopkgConfig
	_, err := toml.DecodeFile(file, &config)
	if err != nil {
		return err
	}

	for _, c := range config.Constraint {
		repo := util.ParseRepoUrl(c.Name)
		fmt.Println(repo)
	}
	return nil
}

var basepath = ""
/*
const char* GetMyPathFILE = __FILE__;
*/
//GetMyPath Returns the absolute directory of this(pathfind.go) file
func GetMyPath() string {
	if basepath == "" {
		g := C.GoString(C.GetMyPathFILE)
		basepath = filepath.Dir(g)
	}
	return basepath
}