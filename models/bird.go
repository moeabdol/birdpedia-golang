package models

import (
	"context"
	"time"
)

// Bird model
type Bird struct {
	ID          int64     `json:"id"`
	Species     string    `json:"species"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const createBirdQuery = `
INSERT INTO birds (
    species,
    description
) VALUES (
    $1, $2
) RETURNING id, species, description, created_at, updated_at;
`

// CreateBirdParams type
type CreateBirdParams struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// CreateBird function
func (store *Store) CreateBird(ctx context.Context, arg CreateBirdParams) (Bird, error) {
	row := store.dbtx.QueryRowContext(ctx, createBirdQuery, arg.Species, arg.Description)
	var bird Bird
	err := row.Scan(
		&bird.ID,
		&bird.Species,
		&bird.Description,
		&bird.CreatedAt,
		&bird.UpdatedAt,
	)
	return bird, err
}

const deleteBirdQuery = `
DELETE FROM birds
WHERE ID = $1;
`

// DeleteBird function
func (store *Store) DeleteBird(ctx context.Context, id int64) error {
	_, err := store.dbtx.ExecContext(context.Background(), deleteBirdQuery, id)
	return err
}

const getBirdQuery = `
SELECT id, species, description, created_at, updated_at
FROM birds
WHERE id = $1
LIMIT 1;
`

// GetBird function
func (store *Store) GetBird(ctx context.Context, id int64) (Bird, error) {
	row := store.dbtx.QueryRowContext(ctx, getBirdQuery, id)
	var bird Bird
	err := row.Scan(
		&bird.ID,
		&bird.Species,
		&bird.Description,
		&bird.CreatedAt,
		&bird.UpdatedAt,
	)
	return bird, err
}

const listBirdsQuery = `
SELECT id, species, description, created_at, updated_at
FROM birds
ORDER BY id
LIMIT $1
OFFSET $2;
`

// ListBirdsParams type
type ListBirdsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

// ListBirds function
func (store *Store) ListBirds(ctx context.Context, arg ListBirdsParams) ([]Bird, error) {
	rows, err := store.dbtx.QueryContext(ctx, listBirdsQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var birds []Bird
	for rows.Next() {
		var bird Bird
		if err := rows.Scan(
			&bird.ID,
			&bird.Species,
			&bird.Description,
			&bird.CreatedAt,
			&bird.UpdatedAt,
		); err != nil {
			return nil, err
		}
		birds = append(birds, bird)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return birds, nil
}
