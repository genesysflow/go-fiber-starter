package router

import (
	"github.com/genesysflow/go-fiber-starter/app/module/article"
	"github.com/genesysflow/go-fiber-starter/internal/inertia"
	"github.com/genesysflow/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	App           fiber.Router
	Cfg           *config.Config
	ArticleRouter *article.ArticleRouter
	Inertia       *inertia.Inertia
}

func NewRouter(fiber *fiber.App, cfg *config.Config, articleRouter *article.ArticleRouter, inertia *inertia.Inertia) *Router {
	return &Router{
		App:           fiber,
		Cfg:           cfg,
		ArticleRouter: articleRouter,
		Inertia:       inertia,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	var m = make(map[string]interface{})
	m["userName"] = "John Doe"

	r.App.Get("/inertia", r.Inertia.Render("Main", m))

	// Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.ArticleRouter.RegisterArticleRoutes()
}
