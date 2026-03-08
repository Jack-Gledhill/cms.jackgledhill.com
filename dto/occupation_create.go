package dto

import (
	"time"

	"github.com/Jack-Gledhill/crud.jackgledhill.com/domain"
)

type OccupationCreate struct {
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	URL      string    `json:"url"`
}

func (d *OccupationCreate) ToEntity() *domain.Occupation {
	return &domain.Occupation{
		Name:     d.Name,
		Position: d.Position,
		Start:    d.Start,
		End:      d.End,
		URL:      d.URL,
	}
}
