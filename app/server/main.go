package main

import (
	"context"
	"fmt"
	"net"
	"todo/app/database"
	"todo/app/rpc"
	"todo/config"
	"todo/utils"

	pb "todo/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	conf := config.ParseConfig()
	logger := utils.InitLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Error().Err(err)
	}

	s := grpc.NewServer()
	pgInstance, err := database.NewPG(context.Background(), conf.DSN)
	if err != nil {
		log.Error().Err(err)
	}
	pb.RegisterTodoServer(s, &rpc.Server{PG: pgInstance, Log: &logger})

	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err)
	}
}
