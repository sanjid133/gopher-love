package cmds

import (
	"bufio"
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
	var file string
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

			} else if file != "" {
				f, err := os.Open(file)
				if err != nil {
					log.Fatalln(err)
				}
				defer f.Close()

				reader := bufio.NewReader(f)
				for {
					line, isPrefix, err := reader.ReadLine()
					if err != nil || isPrefix {
						break
					}

					err = pkg.LoveOrganization(string(line))
					if err != nil {
						log.Fatalln(err)
					}
				}
			} else {
				directory, err = os.Getwd()
				if err != nil {
					fmt.Println(err)
				}
				err := pkg.LoveDependency(directory)
				if err != nil {
					log.Fatalln(err)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&org, "organization", "o", "", "github/kubernetes")
	cmd.Flags().StringVarP(&directory, "repository", "r", "", "github/kubernetes/xyz")
	cmd.Flags().StringVarP(&file, "file", "f", "", "/txt/file/path")
	//	cmd.AddCommand(NewOrgCmd())

	return cmd
}
