package controller

import (
	"git.dev-null.rocks/alexohneander/gosearch/pkg/index"
	"github.com/gofiber/fiber/v2"
)

type Document struct {
	Url     string `json:"url" xml:"url" form:"url"`
	Content string `json:"content" xml:"content" form:"content"`
}

func AddDocumentToIndex(c *fiber.Ctx) error {
	doc := new(Document)

	if err := c.BodyParser(doc); err != nil {
		return err
	}

	index.AddDocToIndex(doc.Url, doc.Content)

	return c.SendString("Document added")
}
