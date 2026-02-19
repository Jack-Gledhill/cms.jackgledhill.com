package domain

import (
	"errors"
	"time"
)

var (
	ErrHackathonNotFound = errors.New("hackathon not found")
	ErrHackathonExists   = errors.New("hackathon already exists")
)

type Hackathon struct {
	ID         int
	Title      string
	Date       time.Time
	DevpostURL string
	ProjectID  int
}
