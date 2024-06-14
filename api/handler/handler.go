package handler

import pb "github.com/Javokhdev/Auth-Service/service"

type Handler struct {
	User   *pb.UserService
}

func NewHandler(us *pb.UserService) *Handler {
	return &Handler{us}
}
