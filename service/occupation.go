package service

import (
	"context"

	"github.com/Jack-Gledhill/crud.jackgledhill.com/domain"
	"github.com/Jack-Gledhill/crud.jackgledhill.com/repository"
)

type Occupation struct {
	Repo repository.Occupation
}

func (s *Occupation) Create(ctx context.Context, occupation *domain.Occupation) (uint, error) {
	return s.Repo.Create(ctx, occupation)
}

func (s *Occupation) GetByID(ctx context.Context, id uint) (*domain.Occupation, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *Occupation) GetAll(ctx context.Context) ([]*domain.Occupation, error) {
	return s.Repo.FindAll(ctx)
}

func (s *Occupation) Update(ctx context.Context, occupation *domain.Occupation) error {
	return s.Repo.Update(ctx, occupation)
}

func (s *Occupation) Delete(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}
