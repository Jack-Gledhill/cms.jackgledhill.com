package dto

import "time"

type HackathonResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	DevpostURL string    `json:"devpost_url"`
	ProjectID  int       `json:"project_id"`
}
