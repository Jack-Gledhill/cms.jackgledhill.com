package dto

type HackathonCreate struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	DevpostURL string `json:"devpost_url"`
	ProjectID  uint   `json:"project_id"`
}
