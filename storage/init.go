package postgres

import (
	pb "github.com/Javokhdev/Auth-Service/genprotos"
)

type InitRoot interface {
	User() User
}
type User interface {
	CreateUser(user *pb.Users) (*pb.Void, error)
	GetByIdUser(id *pb.ById) (*pb.Users, error)
	GetAllUser(_ *pb.Users) (*pb.GetAllUsers, error)
	UpdateUser(user *pb.Users) (*pb.Void, error)
	DeleteUser(id *pb.ById) (*pb.Void, error)
	LoginUser(user *pb.Users ) (*pb.Users, error)
}
