package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/messx/go-grpc-playground/grpcscorer/grpcscorerprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func GetCurrentScore() error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScorerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetScore(ctx, &pb.ScoreRequest{MatchId: "1"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Score is %v", r.CurrentScore)
	return nil
}

func StreamCurrentScore() error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScorerClient(conn)

	stream, err := c.StreamScore(context.Background(), &pb.ScoreRequest{MatchId: "1"})

	if err != nil {
		log.Fatalf("could not fetch score: %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Resp received: %s", resp.CurrentScore)
		}
	}()
	<-done
	log.Print("finished")
	return nil
}

func main() {
	flag.Parse()
	StreamCurrentScore()

}
