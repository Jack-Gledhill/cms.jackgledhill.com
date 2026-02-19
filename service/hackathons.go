package service

import (
	"context"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
	"github.com/Jack-Gledhill/cms.jackgledhill.com/repository"
)

type HackathonService struct {
	Repo repository.HackathonRepository
}

func (s *HackathonService) Create(ctx context.Context, hackathon *domain.Hackathon) error {
	return s.Repo.Create(ctx, hackathon)
}

func (s *HackathonService) GetByID(ctx context.Context, id int) (*domain.Hackathon, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *HackathonService) GetAll(ctx context.Context) ([]*domain.Hackathon, error) {
	return s.Repo.FindAll(ctx)
}

func (s *HackathonService) Update(ctx context.Context, hackathon *domain.Hackathon) error {
	return s.Repo.Update(ctx, hackathon)
}

func (s *HackathonService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
