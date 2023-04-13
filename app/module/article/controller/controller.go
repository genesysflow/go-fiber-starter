package controller

import "github.com/genesysflow/go-fiber-starter/app/module/article/service"

type Controller struct {
	Article ArticleController
}

func NewController(articleService service.ArticleService) *Controller {
	return &Controller{
		Article: NewArticleController(articleService),
	}
}
