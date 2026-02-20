package dto

import "time"

type HackathonResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	DevpostURL string    `json:"devpost_url"`
	ProjectID  uint      `json:"project_id"`
}
