package handler

import pb "github.com/Javokhdev/Auth-Service/genprotos"

type Handler struct {
	User   pb.UserServiceClient
}

func NewHandler(us pb.UserServiceClient) *Handler {
	return &Handler{us}
}
