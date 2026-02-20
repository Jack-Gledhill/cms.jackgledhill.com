package domain

import "time"

type Hackathon struct {
	ID         uint
	Title      string
	Date       time.Time
	DevpostURL string
	ProjectID  int
}
