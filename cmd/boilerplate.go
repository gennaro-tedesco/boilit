package cmd

import (
	"fmt"
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

func createDocs(pluginPath string, pluginName string) {
	docsPath := filepath.Join(pluginPath, "doc")
	err := os.MkdirAll(docsPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	tagsPath := filepath.Join(docsPath, "tags")
	tagsFile, tagsErr := os.Create(tagsPath)
	if tagsErr != nil {
		log.Fatal(tagsErr)
	}
	defer tagsFile.Close()

	helpPath := filepath.Join(docsPath, pluginName+".txt")
	helpFile, helpErr := os.Create(helpPath)
	if helpErr != nil {
		log.Fatal(helpErr)
	}
	defer helpFile.Close()
}

func createVimDir(pluginPath string, pluginName string) {
	vimPluginPath := filepath.Join(pluginPath, "plugin")
	fmt.Println(vimPluginPath)
	err := os.MkdirAll(vimPluginPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	pluginVim := filepath.Join(vimPluginPath, pluginName+".vim")
	pluginVimFile, pluginVimErr := os.Create(pluginVim)
	if pluginVimErr != nil {
		log.Fatal(pluginVimErr)
	}
	defer pluginVimFile.Close()

	reloadVim := filepath.Join(vimPluginPath, "reload.vim")
	reloadVimFile, reloadVimErr := os.Create(reloadVim)
	if reloadVimErr != nil {
		log.Fatal(reloadVimErr)
	}
	defer reloadVimFile.Close()

}

func boilPlugin(rootPath string, pluginName string) {
	pluginPath := createPluginDir(rootPath, pluginName)
	createReadme(pluginPath, pluginName)
	createGitignore(pluginPath)
	createDocs(pluginPath, pluginName)
	createVimDir(pluginPath, pluginName)
}
