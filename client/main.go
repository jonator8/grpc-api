package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jonator8/grpc-api/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewNewsApiClient(conn)

	printNews(client)
}

func printNews(client pb.NewsApiClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	news, err := client.GetNews(ctx, &pb.GetNewsRequest{})
	if err != nil {
		log.Fatalf("client.GetNews failed: %v", err)
	}

	fmt.Println(news)
}
