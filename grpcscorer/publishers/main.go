package main

/*
ToDo:
Add a module to get the score for specific match id which then publish the score in Redis pubsub.
Write generic code to subscibe for match scores



*/

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	pb "github.com/messx/go-grpc-playground/grpcscorer/grpcscorerprotos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedScorerServer
}

func (s *server) GetRandomScore() []string {
	scores := make([]string, 10)
	for i, _ := range scores {
		scores[i] = "score is " + strconv.Itoa(100+10*i) + "/1"

	}
	return scores
}

func (s *server) GetScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	log.Printf("Recieved request for match : %v", in.MatchId)
	return &pb.ScoreResponse{MatchId: in.MatchId, CurrentScore: "100/2"}, nil
}

func (s *server) StreamScore(in *pb.ScoreRequest, stream pb.Scorer_StreamScoreServer) error {
	log.Printf("Fetching score for match: %v", in.MatchId)

	scores := s.GetRandomScore()
	var wg sync.WaitGroup
	for i, score := range scores {
		wg.Add(1)
		go func(score string, i int, matchId string) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			resp := &pb.ScoreResponse{MatchId: matchId, CurrentScore: score}
			if err := stream.Send(resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("Finished score %d", i)
		}(score, int(i), in.MatchId)
	}

	wg.Wait()
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScorerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
