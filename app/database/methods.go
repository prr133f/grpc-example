package database

import (
	"context"
	"todo/app/models"

	"github.com/google/uuid"
)

func (p *Postgres) CreateTask(title, description string) (uuid.UUID, error) {
	var id uuid.UUID
	if err := p.DB.QueryRow(context.Background(), `
	INSERT INTO task_schema.task(title, description)
	VALUES($1, $2)
	RETURNING id
	`, title, description).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (p *Postgres) ResolveTask(id uuid.UUID) error {
	if err := p.DB.QueryRow(context.Background(), `
	UPDATE task_schema 
	SET resolved = true
	WHERE id = $1
	`, id).Scan(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetTaskList() ([]models.Task, error) {
	var tasks []models.Task

	rows, err := p.DB.Query(context.Background(), `
	SELECT id, title, description. resolved
	FROM task_schema.task`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task

		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Resolved)

		tasks = append(tasks, task)
	}

	return tasks, nil
}
