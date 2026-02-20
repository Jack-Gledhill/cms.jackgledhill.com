package repository

import (
	"context"
	"errors"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

var (
	ErrOccupationNotFound = errors.New("occupation not found")
	ErrOccupationExists   = errors.New("occupation already exists")
)

type OccupationRepository interface {
	Create(ctx context.Context, occupation *domain.Occupation) error
	FindByID(ctx context.Context, id uint) (*domain.Occupation, error)
	FindAll(ctx context.Context) ([]*domain.Occupation, error)
	Update(ctx context.Context, occupation *domain.Occupation) error
	Delete(ctx context.Context, id uint) error
}
