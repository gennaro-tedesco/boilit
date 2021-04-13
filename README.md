<h1 align="center">
  boilit
</h1>

<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/zathuraconf/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/boilit"/>
  </a>
</h2>

<h4 align="center">Boil yourself a sweet plugin</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a>
</h3>

Ain't nobody got time to create plugin directories: `boilit` yourself! `boilit` creates boilerplate directory structure and code files for neovim plugins; although there is no unique way to do so, we create a reasonable tree structure useful for most use cases, so that you can build on it.


## Installation
Go get it!
```
go get -u -v github.com/gennaro-tedesco/boilit
```

## Usage
All you have to do is thinking of an awesome name for your plugin: once you have it
```
boilit nvim-awesome-plugin
```
creates a plugin boilerplate structure as
```
├── doc
│  ├── nvim-awesome-plugin.txt
│  └── tags
├── lua
│  └── nvim-awesome-plugin
│     ├── config.lua
│     ├── init.lua
│     └── main.lua
├── plugin
│  ├── nvim-awesome-plugin.vim
│  └── reload.vim
└── README.md
```
Watch it in action:

[![asciicast](https://asciinema.org/a/VpggIG2YeksuuryIHFmVATX43.svg)](https://asciinema.org/a/VpggIG2YeksuuryIHFmVATX43)

The plugin skeleton is created by default in the user's home directory `~`: you can specify a custom location via the `-p` flag
```
boilit nvim-awesome-plugin -p ~/custom/path
```
Check the help `boilit -h`.

Batteries included:

- headers and description of what goes in what file
- relative imports of lua modules
- a useful `reload.vim` function to reload your changes without having to exit and reload neovim

## Feedback
If you find this plugin useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!


