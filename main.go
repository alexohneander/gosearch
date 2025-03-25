package main

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/http"
	"git.dev-null.rocks/alexohneander/gosearch/pkg/index"
)

func main() {
	index.TestIndex()

	// Start HTTP Server
	http.StartService()
}
