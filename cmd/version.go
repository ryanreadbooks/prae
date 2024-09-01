package cmd

import (
	"fmt"
	"runtime"
	"strings"

	inver "github.com/ryanreadbooks/prae/internal/ver"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the prae version",
	Long:  "Print the prae version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("prae version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func defineVersion(ver string) {
	if ver == "" {
		version = "unknown"
		return
	}

	v := strings.Split(ver, "\n")[0]
	if !strings.HasPrefix(v, "v") {
		version = "unknown"
		return
	}

	version = v
	inver.SetVersion(version)
}
