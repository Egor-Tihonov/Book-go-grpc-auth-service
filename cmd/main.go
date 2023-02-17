package main

import (
	"net"
	"time"

	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/config"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/db"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/handlers"
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/auth"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/services"
	userSe "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/user"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	svc := userSe.RegisterHandlers(c)

	if err != nil {
		logrus.Fatalf("error load configs: %w", err)
	}

	dbP, err := db.New(c.DBUrl)
	if err != nil {
		logrus.Fatalf("error connecting to db, %w", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		logrus.Fatalln("Failed to listing:", err)
	}

	logrus.Info("------ START SERVER ON ", c.Port, " ------")

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		ExpirationHours: time.Now().Add(1 * time.Hour),
	}

	s := services.New(dbP, &jwt, svc)
	h := handlers.New(s)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}
