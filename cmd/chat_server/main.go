package main

import (
	chatAPI "chat-server/pkg/chat_v1"
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 8004

type server struct {
	chatAPI.UnimplementedChatAPIServer
}

func (s *server) Create(ctx context.Context, req *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	log.Println(req.Usernames)
	return &chatAPI.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}
func (s *server) Delete(ctx context.Context, req *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Println(req.Id)
	return &emptypb.Empty{}, nil
}
func (s *server) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println(req.Text)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal("Failed to connect to the port")
	}
	s := grpc.NewServer()
	reflection.Register(s)
	chatAPI.RegisterChatAPIServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatal("Server fell down while serving...\n", err)
	}
}
