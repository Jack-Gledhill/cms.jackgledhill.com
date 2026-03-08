package dto

import "github.com/Jack-Gledhill/cms.jackgledhill.com/domain"

type OccupationList struct {
	Occupation []OccupationResponse `json:"occupation"`
}

func (d OccupationList) FromEntities(os []*domain.Occupation) {
	for _, o := range os {
		d.Occupation = append(d.Occupation, OccupationResponse{
			ID:       o.ID,
			Name:     o.Name,
			Position: o.Position,
			Start:    o.Start,
			End:      o.End,
			URL:      o.URL,
		})
	}
}
