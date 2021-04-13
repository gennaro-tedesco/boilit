package cmd

import (
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
		boilPlugin(rootPath, args[0])
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
	rootCmd.Flags().StringP("path", "p", homedir, "root path of plugin directory")
}

func getRootHelp() string {
	return `
boilit: create boilterplate code for neovim plugins

Arguments:
  plugin-name    name of the plugin to create

Usage:
  boilit plugin-name [flag]

Flags:
  -h, --help   help for boilit
  -p, --path   root path of plugin directory: defaults to user's home directory: ~
`
}
