package rpc

import (
	"todo/app/models"

	pb "todo/proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Server struct {
	pb.UnimplementedTodoServer
	PG  RPCIface
	Log *zerolog.Logger
}

type RPCIface interface {
	CreateTask(title, description string) (uuid.UUID, error)
	ResolveTask(id uuid.UUID) error
	GetTaskList() ([]models.Task, error)
	GetTaskById(id uuid.UUID) (models.Task, error)
	DeleteTask(id uuid.UUID) error
}
