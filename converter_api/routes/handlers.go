package routes

import (
	"github.com/andreabreu76/converter_api/controllers"
	"github.com/andreabreu76/converter_api/middlewares"
	"github.com/gorilla/mux"
)

type ExchangeRoutes struct {
	handlers *controllers.ExchangeHandlers
}

func NewExchangeRoutes(handlers *controllers.ExchangeHandlers) *ExchangeRoutes {
	return &ExchangeRoutes{
		handlers: handlers,
	}
}

func (r *ExchangeRoutes) SetupRoutes(router *mux.Router) {

	router.HandleFunc("/exchange/{amount}/{from}/{to}/{rate}", r.handlers.ConvertCurrency).Methods("GET")

	secured := router.PathPrefix("/exchange").Subrouter()
	secured.Use(middlewares.AuthMiddleware)
	secured.HandleFunc("/", r.handlers.CreateExchange).Methods("POST")
	secured.HandleFunc("/{id}", r.handlers.GetExchange).Methods("GET")
	secured.HandleFunc("/{id}", r.handlers.UpdateExchange).Methods("PUT")
	secured.HandleFunc("/{id}", r.handlers.DeleteExchange).Methods("DELETE")
	secured.HandleFunc("/list", r.handlers.ListExchanges).Methods("GET")
}
