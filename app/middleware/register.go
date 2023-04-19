package middleware

import (
	"time"

	"github.com/genesysflow/go-fiber-starter/internal/inertia"
	"github.com/genesysflow/go-fiber-starter/utils"
	"github.com/genesysflow/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Middleware is a struct that contains all the middleware functions
type Middleware struct {
	App     *fiber.App
	Cfg     *config.Config
	Inertia inertia.Inertia
}

func NewMiddleware(app *fiber.App, cfg *config.Config, inertia *inertia.Inertia) *Middleware {
	return &Middleware{
		App:     app,
		Cfg:     cfg,
		Inertia: *inertia,
	}
}

// Register registers all the middleware functions
func (m *Middleware) Register() {
	// Add Extra Middlewares
	m.App.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(m.Cfg.Middleware.Limiter.Enable),
		Max:        m.Cfg.Middleware.Limiter.Max,
		Expiration: m.Cfg.Middleware.Limiter.Expiration * time.Second,
	}))

	m.App.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(m.Cfg.Middleware.Compress.Enable),
		Level: m.Cfg.Middleware.Compress.Level,
	}))

	m.App.Use(recover.New(recover.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Recover.Enable),
	}))

	m.App.Use(pprof.New(pprof.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Pprof.Enable),
	}))

	m.App.Get(m.Cfg.Middleware.Monitor.Path, monitor.New(monitor.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Monitor.Enable),
	}))

	m.App.Use(inertia.New(
		m.Inertia,
	))
}
