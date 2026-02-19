package repository

import (
	"context"
	"database/sql"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type SQLHackathonRepository struct {
	db *sql.DB
}

func NewSQLHackathonRepository(db *sql.DB) *SQLHackathonRepository {
	return &SQLHackathonRepository{
		db: db,
	}
}

func (r *SQLHackathonRepository) Create(ctx context.Context, e *domain.Hackathon) error {
	query := `INSERT INTO hackathons (title, date, devpost_url, project_id) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, e.Title, e.Date, e.DevpostURL, e.ProjectID)
	return err
}

func (r *SQLHackathonRepository) FindByID(ctx context.Context, id int) (*domain.Hackathon, error) {
	query := `SELECT id, title, date, devpost_url, project_id FROM hackathons WHERE id = $1`

	h := &domain.Hackathon{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&h.ID, &h.Title, &h.Date, &h.DevpostURL, &h.ProjectID)
	return h, err
}

func (r *SQLHackathonRepository) FindAll(ctx context.Context) ([]*domain.Hackathon, error) {
	query := `SELECT id, title, date, devpost_url, project_id FROM hackathons`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hs []*domain.Hackathon
	for rows.Next() {
		h := &domain.Hackathon{}
		err = rows.Scan(&h.ID, &h.Title, &h.Date, &h.DevpostURL, &h.ProjectID)
		if err != nil {
			return nil, err
		}

		hs = append(hs, h)
	}

	return hs, nil
}

func (r *SQLHackathonRepository) Update(ctx context.Context, e *domain.Hackathon) error {
	query := `UPDATE hackathons SET title = $1, date = $2, devpost_url = $3, project_id = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query, e.Title, e.Date, e.DevpostURL, e.ProjectID, e.ID)
	return err
}

func (r *SQLHackathonRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM hackathons WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
