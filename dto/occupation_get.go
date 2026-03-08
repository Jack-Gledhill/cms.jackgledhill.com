package dto

import "time"

type OccupationResponse struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	URL      string    `json:"url"`
}
