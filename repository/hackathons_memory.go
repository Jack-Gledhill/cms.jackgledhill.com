package repository

import (
	"context"
	"maps"
	"slices"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type MemoryHackathonRepository struct {
	mu       sync.RWMutex
	entities map[int]*domain.Hackathon
}

func NewMemoryHackathonRepository() *MemoryHackathonRepository {
	return &MemoryHackathonRepository{
		entities: make(map[int]*domain.Hackathon),
	}
}

func (r *MemoryHackathonRepository) Create(_ context.Context, e *domain.Hackathon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[e.ID]; exists {
		return ErrHackathonExists
	}
	r.entities[e.ID] = e
	return nil
}

func (r *MemoryHackathonRepository) FindByID(_ context.Context, id int) (*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	e, exists := r.entities[id]
	if !exists {
		return nil, ErrHackathonNotFound
	}
	return e, nil
}

func (r *MemoryHackathonRepository) FindAll(_ context.Context) ([]*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return slices.Collect(maps.Values(r.entities)), nil
}

func (r *MemoryHackathonRepository) Update(_ context.Context, e *domain.Hackathon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[e.ID]; !exists {
		return ErrHackathonNotFound
	}
	r.entities[e.ID] = e
	return nil
}

func (r *MemoryHackathonRepository) Delete(_ context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[id]; !exists {
		return ErrHackathonNotFound
	}

	delete(r.entities, id)
	return nil
}
