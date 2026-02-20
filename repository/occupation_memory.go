package repository

import (
	"context"
	"maps"
	"slices"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type MemoryOccupationRepository struct {
	mu       sync.RWMutex
	entities map[int]*domain.Occupation
}

func NewMemoryOccupationRepository() *MemoryOccupationRepository {
	return &MemoryOccupationRepository{
		entities: make(map[int]*domain.Occupation),
	}
}

func (r *MemoryOccupationRepository) Create(_ context.Context, e *domain.Occupation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[e.ID]; exists {
		return ErrOccupationExists
	}
	r.entities[e.ID] = e
	return nil
}

func (r *MemoryOccupationRepository) FindByID(_ context.Context, id int) (*domain.Occupation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	e, exists := r.entities[id]
	if !exists {
		return nil, ErrOccupationNotFound
	}
	return e, nil
}

func (r *MemoryOccupationRepository) FindAll(_ context.Context) ([]*domain.Occupation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return slices.Collect(maps.Values(r.entities)), nil
}

func (r *MemoryOccupationRepository) Update(_ context.Context, e *domain.Occupation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[e.ID]; !exists {
		return ErrOccupationNotFound
	}
	r.entities[e.ID] = e
	return nil
}

func (r *MemoryOccupationRepository) Delete(_ context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.entities[id]; !exists {
		return ErrOccupationNotFound
	}

	delete(r.entities, id)
	return nil
}
