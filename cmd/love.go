package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sanjid133/gopher-love/pkg"
	"log"
	"os"
)

func NewCmdLove() *cobra.Command  {
	var org string
	var token string
	cmd := &cobra.Command{
		Use: "love",
		Short: "show love",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if org != ""{
				o, err := pkg.ValidateOrganization(token, org)
				if err != nil {
					log.Fatalln(err)
				}
				err = o.LoveOrganization()
				if err != nil {
					log.Fatalln(err)
				}
			}
		},

	}
	cmd.Flags().StringVarP(&org, "organization", "o", "", "github/kubernetes")
	cmd.Flags().StringVar(&token, "github-token", os.Getenv("GITHUB_TOKEN"),  "github api token")
	cmd.Flags().StringVar(&token, "gitlab-token", os.Getenv("GITLAB_TOKEN"),  "gitlab api token")
//	cmd.AddCommand(NewOrgCmd())

	return cmd
}
