package controller

import (
	"fmt"
	"os"
	"strings"

	"git.dev-null.rocks/alexohneander/gosearch/pkg/index"
	"git.dev-null.rocks/alexohneander/gosearch/pkg/search"
	"github.com/gofiber/fiber/v2"
)

func SearchQuery(c *fiber.Ctx) error {
	query := c.Params("query")
	query = strings.TrimSpace(query)

	terms, queryType := parseQuery(query)
	results := search.Search(terms, queryType, index.Index, index.DocFreq, len(index.Documents))

	var response string

	response = fmt.Sprintf("Search Results (%s query):\n", queryType)
	for _, result := range results {
		response = response + "\n" + fmt.Sprintf("- %s (Score: %.4f)\n", result.Document, result.Score)
	}

	return c.SendString(response)
}

// parseQuery parses the query to determine query type and terms
func parseQuery(query string) ([]string, string) {
	if strings.Contains(query, "AND") {
		return strings.Split(query, " AND "), "AND"
	} else if strings.Contains(query, "OR") {
		return strings.Split(query, " OR "), "OR"
	}
	return strings.Fields(query), "SIMPLE"
}

// phraseMatch checks if all terms appear in the given document in sequence
func phraseMatch(terms []string, doc string) bool {
	// Read the full document content
	content, err := os.ReadFile(doc)
	if err != nil {
		return false
	}
	// Check if the exact phrase (joined terms) is in the document content
	phrase := strings.Join(terms, " ")
	return strings.Contains(strings.ToLower(string(content)), phrase)
}
