package services

import (
	"context"
	"fmt"
	"time"

	"github.com/andreabreu76/converter_api/entities"
	"github.com/andreabreu76/converter_api/repositories"
)

type ExchangeService interface {
	ConvertCurrency(ctx context.Context, amount float64, from string, to string, rate float64) (*entities.Exchange, error)
	CreateExchange(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error)
	GetExchangeByID(ctx context.Context, id string) (*entities.Exchange, error)
	UpdateExchange(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error)
	DeleteExchange(ctx context.Context, id string) error
	ListExchanges(ctx context.Context) ([]*entities.Exchange, error)
}

type exchangeService struct {
	repo repositories.ExchangeRepository
}

func NewExchangeService(repo repositories.ExchangeRepository) ExchangeService {
	return &exchangeService{repo}
}

func (s *exchangeService) ConvertCurrency(ctx context.Context, amount float64, from string, to string, rate float64) (*entities.Exchange, error) {
	convertedValue := amount * rate

	exchange := &entities.Exchange{
		Amount:         amount,
		FromCurrency:   from,
		ToCurrency:     to,
		Rate:           rate,
		ConvertedValue: convertedValue,
		Date:           time.Now(),
	}

	newExchange, err := s.repo.Save(ctx, exchange)
	if err != nil {
		return nil, fmt.Errorf("service convertCurrency error: %v", err)
	}

	return newExchange, nil
}

func (s *exchangeService) CreateExchange(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error) {
	return s.repo.Save(ctx, exchange)
}

func (s *exchangeService) GetExchangeByID(ctx context.Context, id string) (*entities.Exchange, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *exchangeService) UpdateExchange(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error) {
	return s.repo.Update(ctx, exchange)
}

func (s *exchangeService) DeleteExchange(ctx context.Context, id string) error {
	exchange, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("service DeleteExchange error: %v", err)
	}
	return s.repo.Delete(ctx, exchange)
}

func (s *exchangeService) ListExchanges(ctx context.Context) ([]*entities.Exchange, error) {
	return s.repo.GetAll(ctx)
}
