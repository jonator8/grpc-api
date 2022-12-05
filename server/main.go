package server

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/jonator8/grpc-api/protos"
)

var newsData = []byte(`[
	{
		"id": "9E88FFBB-CE5C-49F8-9B1A-4010DD335A20",
		"title": "new1",
		"text": "this is the new 1",
	},
	{
		"id": "DF692EEB-C11C-4BD2-BD84-4040B301D1AB",
		"title": "new2",
		"text": "this is the new 2",
	},
	{
		"id": "F3C453DF-093C-4348-8F2E-68AAAC5F3646",
		"title": "new3",
		"text": "this is the new 3",
	},
]`)

type ApiServer struct {
	pb.UnimplementedNewsApiServer

	savedNews []*pb.New
}

func (s *ApiServer) GetNews(ctx context.Context, req *pb.GetNewsRequest) (*pb.NewsResponse, error) {
	return &pb.NewsResponse{Data: s.savedNews}, nil
}

func (s *ApiServer) GetNew(ctx context.Context, req *pb.GetNewRequest) (*pb.New, error) {
	for _, new := range s.savedNews {
		if new.Id == req.GetId() {
			return &pb.New{
				Id:    new.Id,
				Title: new.Title,
				Text:  new.Text,
			}, nil
		}
	}
	return nil, fmt.Errorf("new with id %s doesn't exists", req.GetId())
}

func newServer() *ApiServer {
	s := &ApiServer{}
	if err := json.Unmarshal(newsData, &s.savedNews); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}

	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNewsApiServer(grpcServer, newServer())

	fmt.Println("server running at localhost:5051")
	grpcServer.Serve(lis)
}
