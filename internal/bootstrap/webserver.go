package bootstrap

import (
	"context"
	"flag"
	"os"
	"runtime"
	"time"

	"github.com/genesysflow/go-fiber-starter/app/middleware"
	"github.com/genesysflow/go-fiber-starter/app/router"
	"github.com/genesysflow/go-fiber-starter/internal/bootstrap/cache"
	"github.com/genesysflow/go-fiber-starter/internal/bootstrap/database"
	"github.com/genesysflow/go-fiber-starter/utils/config"
	"github.com/genesysflow/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// initialize the webserver
func NewFiber(cfg *config.Config) *fiber.App {
	engine := jet.New("./frontend", ".jet")

	// setup
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
		Views:                 engine,
	})

	// pass production config to check it
	response.IsProduction = cfg.App.Production

	app.Static("/assets", "./frontend/dist/assets")

	return app
}

// function to start webserver
func Start(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	fiber *fiber.App,
	router *router.Router,
	middlewares *middleware.Middleware,
	database *database.Database,
	cache *cache.Cache,
	log zerolog.Logger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddress(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// Information message
				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())
					log.Debug().Msgf("Processes: %d", procs)
					log.Debug().Msgf("PID: %d", os.Getpid())
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Debug().Msg("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()

				cache.ConnectCache()

				database.ConnectDatabase()

				migrate := flag.Bool("migrate", false, "migrate the database")
				seeder := flag.Bool("seed", false, "seed the database")
				flag.Parse()

				// read flag -migrate to migrate the database
				if *migrate {
					database.MigrateModels()
				}
				// read flag -seed to seed the database
				if *seeder {
					database.SeedModels()
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutting down the database")
				database.ShutdownDatabase()

				log.Info().Msg("2- Shutting down the cache")
				cache.ShutdownCache()

				log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
