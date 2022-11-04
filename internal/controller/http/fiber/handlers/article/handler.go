package article

import "github.com/gofiber/fiber/v2"

type UseCase interface {
	getArticlesUseCase
}

type handler struct {
	uc UseCase
}

func InitRoutes(router fiber.Router, uc UseCase) {
	h := handler{
		uc: uc,
	}

	routerGroup := router.Group("/articles")
	routerGroup.Get("/", h.getArticles)
}
