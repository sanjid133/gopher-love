package cmds

import (
	"bufio"
	"log"
	"os"

	"github.com/sanjid133/gopher-love/pkg"
	_ "github.com/sanjid133/gopher-love/pkg/manager"
	_ "github.com/sanjid133/gopher-love/pkg/platform"
	"github.com/spf13/cobra"
)

func NewCmdFollow() *cobra.Command {
	var user string
	var file string
	var err error
	cmd := &cobra.Command{
		Use:               "follow",
		Short:             "follow a user",
		Example:           "gopher follow -u github.com/masudur-rahman",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if user != "" {
				err = pkg.FollowUser(user)
				if err != nil {
					log.Fatalln(err)
				}
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

					err = pkg.FollowUser(string(line))
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		},
	}
	cmd.Flags().StringVarP(&user, "user", "u", "", "github/kubernetes")
	cmd.Flags().StringVarP(&file, "file", "f", "", "/txt/file/path")
	//	cmd.AddCommand(NewOrgCmd())

	return cmd
}
