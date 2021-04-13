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
	pluginVimFile.WriteString(`if exists('g:loaded_` + pluginName + `') | finish | endif

" expose vim commands and interface here
" nnoremap <Plug>PlugCommand :lua require(...).plug_command()<CR>

let s:save_cpo = &cpo
set cpo&vim

let g:loaded_` + pluginName + ` = 1

let &cpo = s:save_cpo
unlet s:save_cpo
`)

	reloadVim := filepath.Join(vimPluginPath, "reload.vim")
	reloadVimFile, reloadVimErr := os.Create(reloadVim)
	if reloadVimErr != nil {
		log.Fatal(reloadVimErr)
	}
	defer reloadVimFile.Close()
	reloadVimFile.WriteString(`function! Reload() abort
	lua for k in pairs(package.loaded) do if k:match("^` + pluginName + `") then package.loaded[k] = nil end end
	lua require("` + pluginName + `")
endfunction

nnoremap rr :call Reload()<CR>
`)
}

func createLuaDir(pluginPath string, pluginName string) {
	luaPath := filepath.Join(pluginPath, "lua/"+pluginName)
	err := os.MkdirAll(luaPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	initLua := filepath.Join(luaPath, "init.lua")
	initLuaFile, initErr := os.Create(initLua)
	if initErr != nil {
		log.Fatal(initErr)
	}
	defer initLuaFile.Close()
	initLuaFile.WriteString(`--[[ this module exposes the interface of lua functions:
define here the lua functions that activate the plugin ]]

local main = require("` + pluginName + `.main")
local config = require("` + pluginName + `.config")
`)

	configLua := filepath.Join(luaPath, "config.lua")
	configLuaFile, configErr := os.Create(configLua)
	if configErr != nil {
		log.Fatal(configErr)
	}
	defer configLuaFile.Close()

	mainLua := filepath.Join(luaPath, "main.lua")
	mainLuaFile, mainErr := os.Create(mainLua)
	if mainErr != nil {
		log.Fatal(mainErr)
	}
	defer mainLuaFile.Close()
	mainLuaFile.WriteString(`local config = require("` + pluginName + `.config")`)
}

func boilPlugin(rootPath string, pluginName string) {
	pluginPath := createPluginDir(rootPath, pluginName)
	createReadme(pluginPath, pluginName)
	createGitignore(pluginPath)
	createDocs(pluginPath, pluginName)
	createVimDir(pluginPath, pluginName)
	createLuaDir(pluginPath, pluginName)
	fmt.Printf("boiled %s at %s\n", pluginName, pluginPath)
}
