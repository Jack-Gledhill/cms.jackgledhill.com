package domain

import (
	"strings"
	"time"
)

type Publication struct {
	ID      uint
	Title   string
	Authors []string
	Source  string
	Date    time.Time
	DOI     string
	PDF     string
}

func (p Publication) AuthorsString() string {
	return strings.Join(p.Authors, ", ")
}
