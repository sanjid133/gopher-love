package cmds

import (
	"fmt"
	"github.com/sanjid133/gopher-love/pkg"
	_ "github.com/sanjid133/gopher-love/pkg/manager"
	_ "github.com/sanjid133/gopher-love/pkg/platform"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func NewCmdLove() *cobra.Command {
	var org string
	var directory string
	var err error
	cmd := &cobra.Command{
		Use:               "love",
		Short:             "show love",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if org != "" {
				err = pkg.LoveOrganization(org)
				if err != nil {
					log.Fatalln(err)
				}
			} else if directory != "" {

			} else {
				directory, err = os.Getwd()
				if err != nil {
					fmt.Println(err)
				}
				pkg.LoveDependency(directory)
			}
		},
	}
	cmd.Flags().StringVarP(&org, "organization", "o", "", "github/kubernetes")
	cmd.Flags().StringVarP(&directory, "repository", "r", "", "github/kubernetes/xyz")
	//	cmd.AddCommand(NewOrgCmd())

	return cmd
}
