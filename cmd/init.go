package cmd

import (
	"fmt"

	"github.com/ryanreadbooks/prae/internal/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init appname",
	Short: "Initialize a config for generating code",
	Long: `Initialize (prae init) a config file in yml format which will be used
		when gen command is called.
	`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("prae init needs 1 app name")
		}

		return config.Handle(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
