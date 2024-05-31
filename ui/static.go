package ui

import (
	"embed"
	_ "embed"
)

//go:embed build/*
var UIBuild embed.FS
