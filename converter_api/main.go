package main

import (
	"log"
	"net/http"

	"github.com/andreabreu76/chalenge_convapi/config"
	"github.com/andreabreu76/chalenge_convapi/controllers"
	"github.com/andreabreu76/chalenge_convapi/middlewares"
	"github.com/andreabreu76/chalenge_convapi/repositories"
	"github.com/andreabreu76/chalenge_convapi/routes"
	"github.com/andreabreu76/chalenge_convapi/services"
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
