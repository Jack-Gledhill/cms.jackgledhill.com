package repository

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *domain.Project) error
	FindByID(ctx context.Context, id int) (*domain.Project, error)
	FindAll(ctx context.Context) ([]*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	Delete(ctx context.Context, id int) error
}
