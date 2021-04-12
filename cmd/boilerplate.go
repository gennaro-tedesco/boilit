package cmd

import (
	"os"
	"path/filepath"
)

func createPluginDir(pluginName string, rootPath string) string {
	pluginPath := filepath.Join(rootPath, pluginName)
	os.MkdirAll(pluginPath, os.ModePerm)
	return pluginPath
}
