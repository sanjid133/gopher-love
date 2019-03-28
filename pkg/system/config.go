package system

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/sanjid133/gopher-love/util"
	"io/ioutil"
	"os"
)

const ConfigDir = "/.gopher"
const ConfigFile = "/.gopher/config.yaml"

type SecretConfig struct {
	Github struct {
		ApiToken string `json:"api_token"`
	} `json:"github"`

	GitLab struct {
		ApiKey string `json:"api_key"`
	} `json:"gitlab"`
}

var Config *SecretConfig

func Init() {
	util.EnsureDirectory(util.HomeDirectory() + ConfigDir)
	Config, _ = Initialize()
}

func Initialize() (*SecretConfig, error) {
	config := &SecretConfig{}

	if _, err := os.Stat(util.HomeDirectory() + ConfigFile); err == nil {
		data, err := ioutil.ReadFile(util.HomeDirectory() + ConfigFile)
		if err != nil {
			return nil, err
		}

		jsonData, err := yaml.YAMLToJSON(data)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(jsonData, config); err != nil {
			return nil, err
		}
	} else {
		if err = WriteConfig(config); err != nil {
			return nil, err
		}

	}
	return config, nil
}

func WriteConfig(config *SecretConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	yamlData, err := yaml.JSONToYAML(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(util.HomeDirectory()+ConfigFile, yamlData, 0777)
}
