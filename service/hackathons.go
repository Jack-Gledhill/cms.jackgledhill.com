package service

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
	"github.com/Jack-Gledhill/cms.jackgledhill.com/repository"
)

type Hackathon struct {
	Repo repository.Hackathon
}

func (s *Hackathon) Create(ctx context.Context, hackathon *domain.Hackathon) (uint, error) {
	return s.Repo.Create(ctx, hackathon)
}

func (s *Hackathon) GetByID(ctx context.Context, id uint) (*domain.Hackathon, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *Hackathon) GetAll(ctx context.Context) ([]*domain.Hackathon, error) {
	return s.Repo.FindAll(ctx)
}

func (s *Hackathon) Update(ctx context.Context, hackathon *domain.Hackathon) error {
	return s.Repo.Update(ctx, hackathon)
}

func (s *Hackathon) Delete(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}
