package repository

import (
	"context"
	"errors"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

var (
	ErrHackathonNotFound = errors.New("hackathon not found")
	ErrHackathonExists   = errors.New("hackathon already exists")
)

type HackathonRepository interface {
	Create(ctx context.Context, e *domain.Hackathon) error
	FindByID(ctx context.Context, id uint) (*domain.Hackathon, error)
	FindAll(ctx context.Context) ([]*domain.Hackathon, error)
	Update(ctx context.Context, e *domain.Hackathon) error
	Delete(ctx context.Context, id uint) error
}
