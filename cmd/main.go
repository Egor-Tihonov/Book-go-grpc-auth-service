package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"

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
	InitLog()
	
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
		SecretKey: c.JWTSecretKey,
	}

	s := services.New(dbP, &jwt, svc)
	h := handlers.New(s)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}

func InitLog() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	})

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
