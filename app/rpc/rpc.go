package rpc

import (
	"context"
	"todo/app/database"
	pb "todo/proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Server struct {
	pb.UnimplementedTodoServer
	PG  database.Postgres
	Log *zerolog.Logger
}

func (s *Server) CreateTask(ctx context.Context, in *pb.Task) (*pb.TaskId, error) {
	id, err := s.PG.CreateTask(in.GetTitle(), in.GetDescription())
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	return &pb.TaskId{Id: id.String()}, nil
}

func (s *Server) ResolveTask(ctx context.Context, in *pb.TaskId) (*pb.Ok, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	if err := s.PG.ResolveTask(id); err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	return &pb.Ok{Ok: true}, nil
}

func (s *Server) GetTaskList(in *pb.None, stream pb.Todo_GetTaskListServer) error {
	tasks, err := s.PG.GetTaskList()
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return err
	}

	for _, task := range tasks {
		if err := stream.Send(&pb.Task{
			Id:          task.ID.String(),
			Title:       task.Title,
			Description: task.Description,
			Resolved:    task.Resolved,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetTaskById(ctx context.Context, in *pb.TaskId) (*pb.Task, error) {
	taskId, err := uuid.Parse(in.GetId())
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	task, err := s.PG.GetTaskById(taskId)
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	return &pb.Task{
		Id:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Resolved:    task.Resolved,
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, in *pb.TaskId) (*pb.Ok, error) {
	taskId, err := uuid.Parse(in.GetId())
	if err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	if err := s.PG.DeleteTask(taskId); err != nil {
		s.Log.Error().Err(err).Stack()
		return nil, err
	}

	return &pb.Ok{
		Ok: true,
	}, nil
}
