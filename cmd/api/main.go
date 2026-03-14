package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	fiberhandler "github.com/aryonuwi/tablink-test-BE/internal/delivery/http"
	"github.com/aryonuwi/tablink-test-BE/internal/pkg/config"
	"github.com/aryonuwi/tablink-test-BE/internal/pkg/database"
	appvalidator "github.com/aryonuwi/tablink-test-BE/internal/pkg/validator"
	"github.com/aryonuwi/tablink-test-BE/internal/repository/postgres"
	"github.com/aryonuwi/tablink-test-BE/internal/usecase"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db, err := database.NewPostgresPool(cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	validator := appvalidator.New()

	ingredientRepo := postgres.NewIngredientRepository(db)
	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepo)
	ingredientHandler := fiberhandler.NewIngredientHandler(ingredientUsecase, validator)

	fiberhandler.RegisterRoutes(app, ingredientHandler)

	log.Printf("HTTP server running on :%s", cfg.HTTPPort)
	log.Fatal(app.Listen(":" + cfg.HTTPPort))
}
