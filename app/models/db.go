package models

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Resolved    bool      `db:"resolved"`
}
