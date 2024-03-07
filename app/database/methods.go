package database

import (
	"context"
	"todo/app/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (p *Postgres) CreateTask(title, description string) (uuid.UUID, error) {
	var id uuid.UUID
	if err := p.DB.QueryRow(context.Background(), `
	INSERT INTO task_schema.task(title, description)
	VALUES($1, $2)
	RETURNING id
	`, title, description).Scan(&id); err != nil {
		p.Log.Error().Err(err).Stack()
		return uuid.Nil, errors.WithStack(err)
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
		p.Log.Error().Err(err).Stack()
		return errors.WithStack(err)
	}

	return nil
}

func (p *Postgres) GetTaskList() ([]models.Task, error) {
	rows, err := p.DB.Query(context.Background(), `
	SELECT id, title, description, resolved
	FROM task_schema.task`)
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return nil, errors.WithStack(err)
	}

	tasks, err := pgx.CollectRows[models.Task](rows, pgx.RowToStructByName[models.Task])
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}

func (p *Postgres) GetTaskById(id uuid.UUID) (models.Task, error) {
	row, err := p.DB.Query(context.Background(), `
	SELECT id, title, description, resolved
	FROM task_schema.task
	WHERE id=$1`, id)
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return models.Task{}, errors.WithStack(err)
	}

	task, err := pgx.CollectOneRow[models.Task](row, pgx.RowToStructByName[models.Task])
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return models.Task{}, errors.WithStack(err)
	}

	return task, nil
}

func (p *Postgres) DeleteTask(id uuid.UUID) error {
	if _, err := p.DB.Exec(context.Background(), `
	DELETE FROM task_schema.task
	WHERE id=$1`, id); err != nil {
		p.Log.Error().Err(err).Stack()
		return errors.WithStack(err)
	}

	return nil
}
