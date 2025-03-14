package cmd

import (
	"fmt"

	"log"

	"github.com/dragonsecurity/ptaas-api/internal/config"
	"github.com/dragonsecurity/ptaas-api/internal/core/handler"
	"github.com/dragonsecurity/ptaas-api/internal/core/worker"
	"github.com/dragonsecurity/ptaas-api/pkg/client"
	"github.com/dragonsecurity/ptaas-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// Core is the processing logic of the apt
type Core struct {
	Cfg config.Config
	Db  *gorm.DB
}

func (c Core) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "core",
		Short: "start apt core system",
		Run: func(_ *cobra.Command, _ []string) {
			c.main()
		},
	}
}

func (c Core) main() {
	// create new fiber app
	app := fiber.New()

	// create new models interface
	modelsInstance := models.New(c.Db)

	// create pool instance
	pool := worker.New(c.Cfg, client.NewClient(), modelsInstance, c.Cfg.Core.Workers)
	pool.Register()

	// register core handler
	h := handler.Handler{
		Secret:     c.Cfg.Core.Secret,
		WorkerPool: pool,
		DB:         modelsInstance,
	}

	h.Register(app)

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", c.Cfg.Core.Port)); err != nil {
		log.Fatalf("[core] failed to start core system")
	}
}
