package cmds

import (
	"flag"
	"github.com/sanjid133/gopher-love/pkg/system"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gopher",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "gopher",
		Short:              "Give thanks to your fellow library",
		Long:               `Send love to your fellow golang library by giving Github *`,
		DisableAutoGenTag:  true,
		DisableFlagParsing: true,
	}
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})
	system.Init()

	cmd.AddCommand(NewCmdLove())
	cmd.AddCommand(NewCmdConfig())
	return cmd
}
