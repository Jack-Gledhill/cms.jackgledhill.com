package repository

import (
	"context"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type MemoryOccupationRepository struct {
	mu       sync.RWMutex
	entities []*domain.Occupation
}

func NewMemoryOccupationRepository() *MemoryOccupationRepository {
	return &MemoryOccupationRepository{}
}

func (r *MemoryOccupationRepository) Create(_ context.Context, e *domain.Occupation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	e.ID = uint(len(r.entities))
	r.entities = append(r.entities, e)
	return nil
}

func (r *MemoryOccupationRepository) FindByID(_ context.Context, id uint) (*domain.Occupation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if id >= uint(len(r.entities)) {
		return nil, ErrOccupationNotFound
	}

	e := r.entities[id]
	if e == nil {
		return nil, ErrOccupationExists
	}
	return e, nil
}

func (r *MemoryOccupationRepository) FindAll(_ context.Context) ([]*domain.Occupation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var es []*domain.Occupation
	for _, e := range r.entities {
		if e != nil {
			es = append(es, e)
		}
	}
	return es, nil
}

func (r *MemoryOccupationRepository) Update(_ context.Context, e *domain.Occupation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if e.ID >= uint(len(r.entities)) {
		return ErrHackathonNotFound
	}

	if r.entities[e.ID] == nil {
		return ErrHackathonNotFound
	}
	r.entities[e.ID] = e
	return nil
}

func (r *MemoryOccupationRepository) Delete(_ context.Context, id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if id >= uint(len(r.entities)) {
		return ErrHackathonNotFound
	}

	if r.entities[id] == nil {
		return ErrHackathonNotFound
	}
	r.entities[id] = nil
	return nil
}
