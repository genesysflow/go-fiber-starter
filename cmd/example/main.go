package main

import (
	"go.uber.org/fx"

	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/genesysflow/go-fiber-starter/app/middleware"
	"github.com/genesysflow/go-fiber-starter/app/module/article"
	"github.com/genesysflow/go-fiber-starter/app/router"
	"github.com/genesysflow/go-fiber-starter/internal/bootstrap"
	"github.com/genesysflow/go-fiber-starter/internal/bootstrap/database"
	"github.com/genesysflow/go-fiber-starter/utils/config"
	_ "go.uber.org/automaxprocs"
)

// @title                       Go Fiber Starter API Documentation
// @version                     1.0
// @description                 This is a sample API documentation.
// @termsOfService              http://swagger.io/terms/
// @contact.name                Developer
// @contact.email               bangadam.dev@gmail.com
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host                        localhost:8080
// @schemes                     http https
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 "Type 'Bearer {TOKEN}' to correctly set the API Key"
// @BasePath                    /
func main() {
	fx.New(
		/* provide patterns */
		// config
		fx.Provide(config.NewConfig),
		// logging
		fx.Provide(bootstrap.NewLogger),
		// fiber
		fx.Provide(bootstrap.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules
		article.NewArticleModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
