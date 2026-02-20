package repository

import (
	"context"
	"database/sql"
	"errors"

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

	e := &domain.Hackathon{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&e.ID, &e.Title, &e.Date, &e.DevpostURL, &e.ProjectID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrHackathonNotFound
	}

	return e, err
}

func (r *SQLHackathonRepository) FindAll(ctx context.Context) ([]*domain.Hackathon, error) {
	query := `SELECT id, title, date, devpost_url, project_id FROM hackathons`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var es []*domain.Hackathon
	for rows.Next() {
		e := &domain.Hackathon{}
		err = rows.Scan(&e.ID, &e.Title, &e.Date, &e.DevpostURL, &e.ProjectID)
		if err != nil {
			return nil, err
		}

		es = append(es, e)
	}

	return es, rows.Err()
}

func (r *SQLHackathonRepository) Update(ctx context.Context, e *domain.Hackathon) error {
	query := `UPDATE hackathons SET title = $1, date = $2, devpost_url = $3, project_id = $4 WHERE id = $5`

	res, err := r.db.ExecContext(ctx, query, e.Title, e.Date, e.DevpostURL, e.ProjectID, e.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrHackathonNotFound
	}

	return nil
}

func (r *SQLHackathonRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM hackathons WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrHackathonNotFound
	}

	return err
}
