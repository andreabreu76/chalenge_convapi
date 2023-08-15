package main

import (
	"log"
	"net/http"

	"github.com/andreabreu76/converter_api/config"
	"github.com/andreabreu76/converter_api/controllers"
	"github.com/andreabreu76/converter_api/middlewares"
	"github.com/andreabreu76/converter_api/repositories"
	"github.com/andreabreu76/converter_api/routes"
	"github.com/andreabreu76/converter_api/services"
	"github.com/gorilla/mux"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Could not setup database: %v", err)
	}

	exchangeRepo := repositories.NewExchangeRepository(db)
	exchangeService := services.NewExchangeService(exchangeRepo)
	exchangeHandlers := controllers.NewExchangeHandlers(exchangeService)

	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.CORSMiddleware)
	exchangeRoutes := routes.NewExchangeRoutes(exchangeHandlers)
	exchangeRoutes.SetupRoutes(r)

	log.Println("Starting server on :8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
