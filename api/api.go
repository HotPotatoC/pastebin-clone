package api

import (
	"os"

	"github.com/HotPotatoC/pastebin-clone/backend"
)

type Dependency struct {
	Backend backend.Dependency
}

func BaseURL() string {
	if os.Getenv("ENVIRONMENT") == "production" {
		return "TODO"
	}

	return "http://localhost:" + os.Getenv("PORT")
}
