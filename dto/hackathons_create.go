package dto

type HackathonCreateRequest struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	DevpostURL string `json:"devpost_url"`
	ProjectID  int    `json:"project_id"`
}
