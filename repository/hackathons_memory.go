package repository

import (
	"context"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type MemoryHackathonRepository struct {
	mu       sync.RWMutex
	entities []*domain.Hackathon
}

func NewMemoryHackathonRepository() *MemoryHackathonRepository {
	return &MemoryHackathonRepository{}
}

func (r *MemoryHackathonRepository) Create(_ context.Context, e *domain.Hackathon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	e.ID = uint(len(r.entities))
	r.entities = append(r.entities, e)
	return nil
}

func (r *MemoryHackathonRepository) FindByID(_ context.Context, id uint) (*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if id >= uint(len(r.entities)) {
		return nil, ErrHackathonNotFound
	}

	e := r.entities[id]
	if e == nil {
		return nil, ErrHackathonNotFound
	}
	return e, nil
}

func (r *MemoryHackathonRepository) FindAll(_ context.Context) ([]*domain.Hackathon, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var es []*domain.Hackathon
	for _, e := range r.entities {
		if e != nil {
			es = append(es, e)
		}
	}
	return es, nil
}

func (r *MemoryHackathonRepository) Update(_ context.Context, e *domain.Hackathon) error {
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

func (r *MemoryHackathonRepository) Delete(_ context.Context, id uint) error {
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
