package dto

import "github.com/Jack-Gledhill/cms.jackgledhill.com/domain"

type HackathonList struct {
	Hackathons []HackathonResponse `json:"hackathons"`
}

func (d HackathonList) FromEntities(hs []*domain.Hackathon) {
	for _, h := range hs {
		d.Hackathons = append(d.Hackathons, HackathonResponse{
			ID:         h.ID,
			Title:      h.Title,
			Date:       h.Date,
			DevpostURL: h.DevpostURL,
			ProjectID:  h.ProjectID,
		})
	}
}
