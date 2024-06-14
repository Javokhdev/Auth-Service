package main

import (
	"fmt"
	"log"

	"github.com/Javokhdev/Auth-Service/api"
	"github.com/Javokhdev/Auth-Service/api/handler"
	pb "github.com/Javokhdev/Auth-Service/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	skill := pb.NewUserServiceClient(UserConn)
	h := handler.NewHandler(skill)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8081")
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
