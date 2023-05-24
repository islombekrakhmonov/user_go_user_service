package main

import (
	"net"
	"user_service/config"
	"user_service/grpc"
	"user_service/grpc/client"
	"user_service/pkg/logger"
	"user_service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	// ----------------------------------------------
	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)

	}

	log := logger.NewLogger("app", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	// ----------------------------------------------

	store, err := postgres.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Panic("Error connect to postgresql: ", logger.Error(err))
		return
	}
	defer store.CloseDB()

	svcs, err := client.NewGrpcClients(cfg)
  if err != nil {
    log.Panic("client.NewGrpcClients", logger.Error(err))
  }

  grpcServer := grpc.SetUpServer(cfg, log, store, svcs)

  lis, err := net.Listen("tcp", cfg.ServicePort)
  if err != nil {
    log.Panic("net.Listen", logger.Error(err))
  }

  log.Info("GRPC: Server being started...", logger.String("port", cfg.ServicePort))

  if err := grpcServer.Serve(lis); err != nil {
    log.Panic("grpcServer.Serve", logger.Error(err))
  }
}
