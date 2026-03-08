package dto

import (
	"time"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type HackathonCreate struct {
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	DevpostURL string    `json:"devpost_url"`
	ProjectID  uint      `json:"project_id"`
}

func (d *HackathonCreate) ToEntity() *domain.Hackathon {
	return &domain.Hackathon{
		Title:      d.Title,
		Date:       d.Date,
		DevpostURL: d.DevpostURL,
		ProjectID:  d.ProjectID,
	}
}
