package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andreabreu76/converter_api/entities"
	"github.com/andreabreu76/converter_api/services"
	"github.com/andreabreu76/converter_api/utils"
	"github.com/gorilla/mux"
)

type ExchangeHandlers struct {
	service services.ExchangeService
}

func NewExchangeHandlers(service services.ExchangeService) *ExchangeHandlers {
	return &ExchangeHandlers{service}
}

func (h *ExchangeHandlers) CreateExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exchange := &entities.Exchange{}
	if err := json.NewDecoder(r.Body).Decode(exchange); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newExchange, err := h.service.CreateExchange(ctx, exchange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(newExchange); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExchangeHandlers) GetExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	exchange, err := h.service.GetExchangeByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(exchange); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExchangeHandlers) UpdateExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exchange := &entities.Exchange{}
	if err := json.NewDecoder(r.Body).Decode(exchange); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedExchange, err := h.service.UpdateExchange(ctx, exchange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(updatedExchange); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExchangeHandlers) DeleteExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	err := h.service.DeleteExchange(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ExchangeHandlers) ListExchanges(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exchanges, err := h.service.ListExchanges(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(exchanges); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExchangeHandlers) ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	response := entities.ConversionResult{}

	amount, err := strconv.ParseFloat(vars["amount"], 64)
	if err != nil {
		http.Error(w, "Invalid amount value", http.StatusBadRequest)
		return
	}

	rate, err := strconv.ParseFloat(vars["rate"], 64)
	if err != nil {
		http.Error(w, "Invalid rate value", http.StatusBadRequest)
		return
	}

	from := vars["from"]
	to := vars["to"]

	result, err := h.service.ConvertCurrency(ctx, amount, from, to, rate)
	if err != nil {
		http.Error(w, "Error converting currency", http.StatusInternalServerError)
		return
	}

	response.ConvertedValue = result.ConvertedValue
	response.ExchangeSymbol = utils.GetCurrencySymbol(to)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
