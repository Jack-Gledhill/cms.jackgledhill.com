package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/domain"
)

type SQLOccupationRepository struct {
	db *sql.DB
}

func NewSQLOccupationRepository(db *sql.DB) *SQLOccupationRepository {
	return &SQLOccupationRepository{
		db: db,
	}
}

func (r *SQLOccupationRepository) Create(ctx context.Context, e *domain.Occupation) (uint, error) {
	query := `INSERT INTO occupation (name, position, start, end, url) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	res, err := r.db.ExecContext(ctx, query, e.Name, e.Position, e.Start, e.End, e.URL)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (r *SQLOccupationRepository) FindByID(ctx context.Context, id uint) (*domain.Occupation, error) {
	query := `SELECT id, name, position, start, end, url FROM occupation WHERE id = $1`

	e := &domain.Occupation{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&e.ID, &e.Name, &e.Position, &e.Start, &e.End, &e.URL)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrOccupationNotFound
	}

	return e, err
}

func (r *SQLOccupationRepository) FindAll(ctx context.Context) ([]*domain.Occupation, error) {
	query := `SELECT id, name, position, start, end, url FROM occupation`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var es []*domain.Occupation
	for rows.Next() {
		e := &domain.Occupation{}
		err = rows.Scan(&e.ID, &e.Name, &e.Position, &e.Start, &e.End, &e.URL)
		if err != nil {
			return nil, err
		}

		es = append(es, e)
	}

	return es, rows.Err()
}

func (r *SQLOccupationRepository) Update(ctx context.Context, e *domain.Occupation) error {
	query := `UPDATE occupation SET name = $1, position = $2, start = $3, end = $4, url = $5 WHERE id = $6`

	res, err := r.db.ExecContext(ctx, query, e.Name, e.Position, e.Start, e.End, e.URL, e.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrOccupationNotFound
	}

	return nil
}

func (r *SQLOccupationRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM occupation WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrOccupationNotFound
	}

	return nil
}
