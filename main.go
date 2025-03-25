package main

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/http"
	"git.dev-null.rocks/alexohneander/gosearch/pkg/search"
)

func main() {
	search.TestIndex()

	// Start HTTP Server
	http.StartService()
}
