package rpc

import (
	"context"
	"io"
	"net"
	"testing"
	pb "todo/proto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestRPC_CreateTask(t *testing.T) {
	ctx := context.Background()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	mockServer := &Server{PG: new(mock)}

	pb.RegisterTodoServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer server.GracefulStop()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	assert.NoError(t, err)

	client := pb.NewTodoClient(conn)

	response, err := client.CreateTask(ctx, &pb.Task{
		Title:       "test",
		Description: "test",
	})
	assert.NoError(t, err)

	assert.NotNil(t, response)
}

func TestRPC_ResolveTask(t *testing.T) {
	ctx := context.Background()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	mockServer := &Server{PG: new(mock)}

	pb.RegisterTodoServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer server.GracefulStop()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	assert.NoError(t, err)

	client := pb.NewTodoClient(conn)

	response, err := client.ResolveTask(ctx, &pb.TaskId{
		Id: uuid.New().String(),
	})
	assert.NoError(t, err)

	assert.NotNil(t, response)
}

func TestRPC_GetTaskList(t *testing.T) {
	ctx := context.Background()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	mockServer := &Server{PG: new(mock)}

	pb.RegisterTodoServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer server.GracefulStop()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	assert.NoError(t, err)

	client := pb.NewTodoClient(conn)

	stream, err := client.GetTaskList(ctx, &pb.None{})
	assert.NoError(t, err)

	for {
		task, err := stream.Recv()
		if err == io.EOF {
			break
		}

		assert.NoError(t, err, "read from stream")
		assert.NotNil(t, task)
	}
}

func TestRPC_GetTaskById(t *testing.T) {
	ctx := context.Background()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	mockServer := &Server{PG: new(mock)}

	pb.RegisterTodoServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer server.GracefulStop()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	assert.NoError(t, err)

	client := pb.NewTodoClient(conn)

	response, err := client.GetTaskById(ctx, &pb.TaskId{
		Id: uuid.NewString(),
	})
	assert.NoError(t, err)

	assert.NotNil(t, response)
}

func TestRPC_DeleteTask(t *testing.T) {
	ctx := context.Background()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	defer server.Stop()

	mockServer := &Server{PG: new(mock)}

	pb.RegisterTodoServer(server, mockServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer server.GracefulStop()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	assert.NoError(t, err)

	client := pb.NewTodoClient(conn)

	response, err := client.DeleteTask(ctx, &pb.TaskId{
		Id: uuid.NewString(),
	})
	assert.NoError(t, err)

	assert.NotNil(t, response)
}
