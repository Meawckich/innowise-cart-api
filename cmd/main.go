package main

import (
	_ "cart-api/docs"
	"cart-api/internal/config"
	"cart-api/internal/handler"
	"cart-api/internal/pkg/db"
	"cart-api/internal/repository"
	"cart-api/internal/server"
	"cart-api/internal/service"
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

//	@title		Trainee cart-Api
//	@version	1.0

// host localhost:3000
// BasePath /
func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to initialize condig: %v", err)
	}

	dbPool, err := db.InitDb()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		log.Println("Closing database connection pool...")
		if err := dbPool.Close(); err != nil {
			log.Fatalf("Error closing database connection pool: %v", err)
		}
		log.Println("Database connection pool closed.")
	}()

	mux := http.NewServeMux()
	//repo
	itemRepo := repository.NewPostgresItemRepository(dbPool)
	cartRepo := repository.NewPostgresCartRepository(dbPool)

	//service
	cartService := service.NewCartService(cartRepo)
	itemService := service.NewItemService(itemRepo, cartRepo)

	//handler
	cartHandler := handler.NewCartHandler(cartService)
	itemHandler := handler.NewItemHandler(itemService)

	cartHandler.GroupRoutes(mux)
	itemHandler.GroupRoutes(mux)
	handler.GroupSwaggerRoute(mux)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := server.NewServer(&cfg, mux)

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
