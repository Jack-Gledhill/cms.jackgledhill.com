package repository

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type Publication interface {
	Create(ctx context.Context, publication *domain.Publication) (uint, error)
	FindByID(ctx context.Context, id uint) (*domain.Publication, error)
	FindAll(ctx context.Context) ([]*domain.Publication, error)
	Update(ctx context.Context, publication *domain.Publication) error
	Delete(ctx context.Context, id uint) error
}
