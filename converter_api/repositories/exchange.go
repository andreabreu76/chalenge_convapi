package repositories

import (
	"context"
	"fmt"

	"github.com/andreabreu76/chalenge_convapi/entities"
	"gorm.io/gorm"
)

type ExchangeRepository interface {
	Save(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error)
	FindByID(ctx context.Context, id string) (*entities.Exchange, error)
	Update(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error)
	Delete(ctx context.Context, exchange *entities.Exchange) error
	GetAll(ctx context.Context) ([]*entities.Exchange, error)
}

type exchangeRepo struct {
	db *gorm.DB
}

func NewExchangeRepository(db *gorm.DB) ExchangeRepository {
	return &exchangeRepo{db}
}

func (r *exchangeRepo) Save(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error) {
	db := r.db.WithContext(ctx)
	err := db.Create(exchange).Error
	return exchange, err
}

func (r *exchangeRepo) FindByID(ctx context.Context, id string) (*entities.Exchange, error) {
	db := r.db.WithContext(ctx)
	exchange := &entities.Exchange{}
	err := db.First(exchange, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("repository FindByID fail: %v", err)
	}
	return exchange, nil
}

func (r *exchangeRepo) Update(ctx context.Context, exchange *entities.Exchange) (*entities.Exchange, error) {
	db := r.db.WithContext(ctx)
	err := db.Save(exchange).Error
	return exchange, err
}

func (r *exchangeRepo) Delete(ctx context.Context, exchange *entities.Exchange) error {
	db := r.db.WithContext(ctx)
	return db.Delete(exchange).Error
}

func (r *exchangeRepo) GetAll(ctx context.Context) ([]*entities.Exchange, error) {
	var exchanges []*entities.Exchange
	db := r.db.WithContext(ctx)
	err := db.Find(&exchanges).Error
	return exchanges, err
}
