package repository

import (
	"context"
	"sync"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type HackathonMemory struct {
	mu       sync.RWMutex
	entities []*domain.Hackathon
}

func NewHackathonMemory() *HackathonMemory {
	return &HackathonMemory{}
}

func (r *HackathonMemory) Create(_ context.Context, e *domain.Hackathon) (uint, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	e.ID = uint(len(r.entities))
	r.entities = append(r.entities, e)
	return e.ID, nil
}

func (r *HackathonMemory) FindByID(_ context.Context, id uint) (*domain.Hackathon, error) {
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

func (r *HackathonMemory) FindAll(_ context.Context) ([]*domain.Hackathon, error) {
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

func (r *HackathonMemory) Update(_ context.Context, e *domain.Hackathon) error {
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

func (r *HackathonMemory) Delete(_ context.Context, id uint) error {
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
