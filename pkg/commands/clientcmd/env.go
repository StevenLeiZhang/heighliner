package clientcmd

import (
	"github.com/spf13/cobra"
)

var (
	envCmd = &cobra.Command{
		Use:   "env",
		Short: "Manage environments",
		Long:  "",
	}
)

func init() {
	envCmd.AddCommand(envNewCmd)
}
