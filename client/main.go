package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jonator8/grpc-api/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewNewsApiClient(conn)

	printNews(client)
	fmt.Println("-------------------")
	printOneNews(client, "DF692EEB-C11C-4BD2-BD84-4040B301D1AB")
}

func printNews(client pb.NewsApiClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	news, err := client.GetNews(ctx, &pb.GetNewsRequest{})
	if err != nil {
		log.Fatalf("client.GetNews failed: %v", err)
	}

	fmt.Println(news.Data)
}

func printOneNews(client pb.NewsApiClient, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	new, err := client.GetNew(ctx, &pb.GetNewRequest{Id: id})
	if err != nil {
		log.Fatalf("client.GetNews failed: %v", err)
	}

	fmt.Println(new)
}
