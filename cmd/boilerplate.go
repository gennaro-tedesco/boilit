package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func createPluginDir(rootPath string, pluginName string) string {
	pluginPath := filepath.Join(rootPath, pluginName)
	err := os.MkdirAll(pluginPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return pluginPath
}

func createReadme(pluginPath string, pluginName string) {
	readmePath := filepath.Join(pluginPath, "README.md")
	file, err := os.Create(readmePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("## " + pluginName + "\n")
	file.WriteString("This is the README for " + pluginName)
}

func createGitignore(pluginPath string) {
	gitignorePath := filepath.Join(pluginPath, ".gitignore")
	file, err := os.Create(gitignorePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("/plugin/reload.vim")

}

func boilPlugin(rootPath string, pluginName string) {
	pluginPath := createPluginDir(rootPath, pluginName)
	createReadme(pluginPath, pluginName)
	createGitignore(pluginPath)
}
