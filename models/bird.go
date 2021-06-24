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
