package cmds

import (
	"fmt"
	"github.com/sanjid133/gopher-love/pkg/system"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "config",
		Short:             "Configure platform",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			platform, key := ask()
			config, err := system.Initialize()
			if err != nil {
				fmt.Println(err)
			}
			switch platform {
			case "GitHub":
				config.Github.ApiToken = key
				break
			case "GitLab":
				config.GitLab.ApiKey = key
				break
			}

			if err = system.WriteConfig(config); err != nil {
				fmt.Println(err)
			}
		},
	}
	//	cmd.AddCommand(NewOrgCmd())

	return cmd
}

func ask() (string, string) {
	platform := ""
	prompt := &survey.Select{
		Message: "Choose a Platform:",
		Options: []string{"GitHub", "GitLab", "BitBucket"},
	}
	survey.AskOne(prompt, &platform, nil)

	token := ""
	p := &survey.Password{
		Message: "Please enter api-key",
	}
	survey.AskOne(p, &token, nil)
	return platform, token
}
