package handlers

import (
	pb "github.com/Egor-Tihonov/go-grpc-auth-service/pkg/pb/auth"
	"github.com/Egor-Tihonov/go-grpc-auth-service/pkg/services"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	se *services.Server
}

func New(s *services.Server) *Handler {
	return &Handler{
		se: s,
	}
}
