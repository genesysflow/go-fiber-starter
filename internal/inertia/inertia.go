package inertia

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/genesysflow/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	vueglue "github.com/torenware/vite-go"
)

type Inertia struct {
	Filter      func(c *fiber.Ctx) bool // Required
	version     string
	url         string
	ssrURL      string
	sharedProps map[string]interface{}
	glue        *vueglue.VueGlue
}

var ConfigDefault = Inertia{
	Filter:  nil,
	version: "",
	url:     "",
	ssrURL:  "",
}

func configDefault(config ...Inertia) Inertia {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	return cfg
}

func (i *Inertia) isInertiaRequest(c *fiber.Ctx) bool {
	return c.Get("X-Inertia") != ""
}

func NewInertia(cfg *config.Config, log zerolog.Logger) *Inertia {
	var i = Inertia{
		version: "",
		url:     cfg.App.URL,
	}

	env := "development"

	if cfg.App.Production {
		env = "production"
	}

	config := &vueglue.ViteConfig{
		Environment: env,
		AssetsPath:  "dist",
		EntryPoint:  "src/main.js",
		Platform:    "vue",
		FS:          os.DirFS("frontend"),
	}

	glue, err := vueglue.NewVueGlue(config)

	if err != nil {
		log.Panic().Err(err).Msg("Failed to glue vue")
	}

	i.glue = glue

	return &i
}

func (i *Inertia) Render(component string, props map[string]interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		only := make(map[string]string)
		partial := c.Get("X-Inertia-Partial-Data")

		if partial != "" && c.Get("X-Inertia-Partial-Component") == component {
			for _, value := range strings.Split(partial, ",") {
				only[value] = value
			}
		}

		page := &Page{
			Component: component,
			Props:     make(map[string]interface{}),
			URL:       c.Path(),
			Version:   i.version,
		}

		for key, value := range i.sharedProps {
			if _, ok := only[key]; len(only) == 0 || ok {
				page.Props[key] = value
			}
		}

		for key, value := range props {
			if _, ok := only[key]; len(only) == 0 || ok {
				page.Props[key] = value
			}
		}

		if i.isInertiaRequest(c) {
			js, err := json.Marshal(page)
			if err != nil {
				return err
			}

			c.Set("Vary", "Accept")
			c.Set("X-Inertia", "true")
			c.Set("Content-Type", "application/json")

			_, err = c.Write(js)
			if err != nil {
				return err
			}

			return nil
		}

		viewData := make(map[string]interface{})
		viewData["page"] = page
		viewData["glue"] = i.glue

		c.Set("Content-Type", "text/html")

		return c.Render("index", viewData)
	}
}

func (i *Inertia) Share(key string, value interface{}) {
	i.sharedProps[key] = value
}

func New(config Inertia) fiber.Handler {
	// For setting default config
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		if c.Get("X-Inertia") == "" {
			return c.Next()
		}

		if c.Method() == "GET" && c.Get("X-Inertia-Version") != cfg.version {
			c.Set("X-Inertia-Location", cfg.url+c.Path())
			return c.SendStatus(fiber.StatusConflict)
		}

		return c.Next()
	}
}
