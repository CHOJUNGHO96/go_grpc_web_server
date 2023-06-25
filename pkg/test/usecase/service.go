package usecase

import (
	"context"
	"fmt"
	pb "sequence_game_server/api/v1/test"
	server "sequence_game_server/pkg"
)

type TestService struct {
	pb.UnimplementedTestServiceServer
}

func init() {
	testService := &TestService{}
	pb.RegisterTestServiceServer(server.GrpcServer, testService)
}

func (s *TestService) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	testStringData := req.TestStringData
	testIntData := req.TestIntData

	fmt.Println(testStringData)
	fmt.Println(testIntData)

	return &pb.TestResponse{
		Result: "success",
	}, nil
}
