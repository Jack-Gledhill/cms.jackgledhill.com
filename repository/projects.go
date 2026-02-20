package repository

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type Project interface {
	Create(ctx context.Context, project *domain.Project) (uint, error)
	FindByID(ctx context.Context, id uint) (*domain.Project, error)
	FindAll(ctx context.Context) ([]*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id uint) error
}
