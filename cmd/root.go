package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"
)

var usr, _ = user.Current()
var homedir = usr.HomeDir

var rootCmd = &cobra.Command{
	Use:   "boilit",
	Args:  cobra.ExactArgs(1),
	Short: "create boilerplate code for neovim plugins",
	Long:  `create boilerplate code for neovim plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		rootPath, _ := cmd.Flags().GetString("path")
		createPluginDir(args[0], rootPath)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("path", "p", homedir, "root path to plugin directory")
}
