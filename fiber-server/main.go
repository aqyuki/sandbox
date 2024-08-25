package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		log.Infof("Serve request from %s", c.IP())
		return c.
			Status(http.StatusOK).
			SendString("Hello, World!\n")
	})

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		if err := app.Shutdown(); err != nil {
			log.Errorf("failed to shutdown server : %v", err)
			return
		}
		log.Info("server shutdown gracefully")
	}()

	log.Info("server started on :8080")
	if err := app.Listen(":8080"); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Errorf("failed to start server : %v", err)
	}
	wg.Wait()
}
