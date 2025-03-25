package main

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/http"
	"git.dev-null.rocks/alexohneander/gosearch/pkg/index"
)

func main() {
	// Initialize Index
	index.InitIndex("default")

	// Start HTTP Server
	http.StartService()
}
