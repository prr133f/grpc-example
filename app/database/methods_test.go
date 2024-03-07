package database

import (
	"os"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMethods_CreateTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.Nil(t, err, "NewPool()")

	tests := struct {
		id          uuid.UUID
		title       string
		description string
	}{
		id:          uuid.New(),
		title:       "test title",
		description: "test description",
	}

	mock.ExpectQuery(
		regexp.QuoteMeta(`
		INSERT INTO task_schema.task(title, description)
		VALUES($1, $2)
		RETURNING id`),
	).WithArgs(tests.title, tests.description).WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(tests.id.String()))

	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Logger()

	pgMock := Postgres{
		DB:  mock,
		Log: &logger,
	}

	id, err := pgMock.CreateTask(tests.title, tests.description)
	assert.Nil(t, err, "CreateTask()")
	assert.Equal(t, id, tests.id)
}

func TestMethods_ResolveTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.Nil(t, err, "NewPool()")

	tests := struct {
		id uuid.UUID
	}{
		id: uuid.New(),
	}

	mock.ExpectExec(
		regexp.QuoteMeta(`
		UPDATE task_schema.task 
		SET resolved = true
		WHERE id = $1
		ON CONFLICT
		DO NOTHING`),
	).WithArgs(tests.id).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Logger()

	pgMock := Postgres{
		DB:  mock,
		Log: &logger,
	}

	require.NoError(t, pgMock.ResolveTask(tests.id), "ResolveTask()")
}

func TestMethods_GetTaskList(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.Nil(t, err, "NewPool()")

	id := uuid.New()

	mock.ExpectQuery(
		regexp.QuoteMeta(`
		SELECT id, title, description, resolved
		FROM task_schema.task`),
	).WillReturnRows(pgxmock.NewRows([]string{"id", "title", "description", "resolved"}).AddRow(id.String(), "test", "test", false))

	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Logger()

	pgMock := Postgres{
		DB:  mock,
		Log: &logger,
	}

	tasks, err := pgMock.GetTaskList()
	assert.Nil(t, err, "GetTaskList()")
	assert.Equal(t, id, tasks[0].ID)
}

func TestMethods_GetTaskById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.Nil(t, err, "NewPool()")

	id := uuid.New()

	mock.ExpectQuery(
		regexp.QuoteMeta(`
		SELECT id, title, description, resolved
		FROM task_schema.task
		WHERE id=$1`),
	).WithArgs(id).WillReturnRows(pgxmock.NewRows([]string{"id", "title", "description", "resolved"}).AddRow(id.String(), "test", "test", false))

	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Logger()

	pgMock := Postgres{
		DB:  mock,
		Log: &logger,
	}

	task, err := pgMock.GetTaskById(id)
	assert.Nil(t, err, "GetTaskList()")
	assert.Equal(t, id, task.ID)
}

func TestMethods_DeleteTask(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.Nil(t, err, "NewPool()")

	id := uuid.New()

	mock.ExpectExec(
		regexp.QuoteMeta(`
		DELETE FROM task_schema.task
		WHERE id=$1`),
	).WithArgs(id).WillReturnResult(pgxmock.NewResult("DELETE", 1))

	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Logger()

	pgMock := Postgres{
		DB:  mock,
		Log: &logger,
	}

	require.NoError(t, pgMock.DeleteTask(id), "DeleteTask()")
}
