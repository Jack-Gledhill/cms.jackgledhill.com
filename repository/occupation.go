package repository

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type OccupationRepository interface {
	Create(ctx context.Context, occupation *domain.Occupation) error
	FindByID(ctx context.Context, id int) (*domain.Occupation, error)
	FindAll(ctx context.Context) ([]*domain.Occupation, error)
	Update(ctx context.Context, occupation *domain.Occupation) error
	Delete(ctx context.Context, id int) error
}
