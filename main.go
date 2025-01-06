package main

import (
	"embed"
	"github.com/SiTiSem/webdu/actions"
)

//go:embed frontend/dist
var frontendEmbed embed.FS

func main() {
	actions.FrontendFs = frontendEmbed
	actions.HttpRoot(frontendEmbed)
}
