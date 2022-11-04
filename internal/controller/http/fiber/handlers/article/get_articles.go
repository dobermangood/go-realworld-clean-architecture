package article

import "github.com/gofiber/fiber/v2"

type getArticlesUseCase interface {
	GetArticles()
}

func (h *handler) getArticles(c *fiber.Ctx) error {
	h.uc.GetArticles()
	return c.SendString("articles")
}
