package repository

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type HackathonRepository interface {
	Create(ctx context.Context, hackathon *domain.Hackathon) error
	FindByID(ctx context.Context, id int) (*domain.Hackathon, error)
	FindAll(ctx context.Context) ([]*domain.Hackathon, error)
	Update(ctx context.Context, hackathon *domain.Hackathon) error
	Delete(ctx context.Context, id int) error
}
