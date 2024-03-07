package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"todo/app/database"
	"todo/app/rpc"
	"todo/utils"

	pb "todo/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	logger := utils.InitLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Error().Err(err)
	}

	s := grpc.NewServer()
	pgInstance, err := database.NewPG(context.Background(), os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Error().Err(err)
	}
	pb.RegisterTodoServer(s, &rpc.Server{PG: pgInstance, Log: &logger})

	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err)
	}
}
