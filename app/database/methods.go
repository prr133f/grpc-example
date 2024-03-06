package database

import (
	"context"
	"todo/app/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	if _, err := p.DB.Exec(context.Background(), `
	UPDATE task_schema.task 
	SET resolved = true
	WHERE id = $1
	ON CONFLICT
	DO NOTHING
	`, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetTaskList() ([]models.Task, error) {
	var tasks []models.Task

	rows, err := p.DB.Query(context.Background(), `
	SELECT id, title, description, resolved
	FROM task_schema.task`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task, err := pgx.RowToStructByName[models.Task](rows)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p *Postgres) GetTaskById(id uuid.UUID) (models.Task, error) {
	row, err := p.DB.Query(context.Background(), `
	SELECT id, title, description, resolved
	FROM task_schema.task`)
	if err != nil {
		return models.Task{}, err
	}

	task, err := pgx.RowToStructByName[models.Task](row)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (p *Postgres) DeleteTask(id uuid.UUID) error {
	if _, err := p.DB.Exec(context.Background(), `
	DELETE FROM task_schema.task
	WHERE id=$1`, id); err != nil {
		return err
	}

	return nil
}
