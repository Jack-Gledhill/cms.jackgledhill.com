package domain

import "time"

type Hackathon struct {
	ID         int
	Title      string
	Date       time.Time
	DevpostURL string
	ProjectID  int
}
