package repository

import (
	"context"
	"maps"
	"slices"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type MemoryHackathonRepository struct {
	mu         sync.RWMutex
	hackathons map[int]*domain.Hackathon
}

func NewMemoryHackathonRepository() *MemoryHackathonRepository {
	return &MemoryHackathonRepository{
		hackathons: make(map[int]*domain.Hackathon),
	}
}

func (r *MemoryHackathonRepository) Create(_ context.Context, e *domain.Hackathon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hackathons[e.ID]; exists {
		return domain.ErrHackathonExists
	}
	r.hackathons[e.ID] = e
	return nil
}

func (r *MemoryHackathonRepository) FindByID(_ context.Context, id int) (*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	h, exists := r.hackathons[id]
	if !exists {
		return nil, domain.ErrHackathonNotFound
	}
	return h, nil
}

func (r *MemoryHackathonRepository) FindAll(_ context.Context) ([]*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return slices.Collect(maps.Values(r.hackathons)), nil
}

func (r *MemoryHackathonRepository) Update(_ context.Context, e *domain.Hackathon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hackathons[e.ID]; !exists {
		return domain.ErrHackathonNotFound
	}
	r.hackathons[e.ID] = e
	return nil
}

func (r *MemoryHackathonRepository) Delete(_ context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hackathons[id]; !exists {
		return domain.ErrHackathonNotFound
	}

	delete(r.hackathons, id)
	return nil
}
