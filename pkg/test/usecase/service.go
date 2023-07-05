package usecase

import (
	"context"
	pb "sequence_game_server/api/v1/test"
	"sequence_game_server/core/util"
	"sequence_game_server/pkg/test/model"
	repository "sequence_game_server/pkg/test/repository"
	"time"
)

type Response struct {
	TestResponse []model.SelectCompanyDepartment
}

type TestService struct {
	pb.UnimplementedTestServiceServer
	repo *repository.Repository
}

func NewTestService(repo *repository.Repository) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	testStringData := req.TestStringData
	testIntData := req.TestIntData
	util.Logger.Println("testStringData : ", testStringData, "testIntData : ", testIntData)
	// 5초 제한의 새로운 컨텍스트 생성
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// 리소스 반환을 위한 cancel 함수 호출 지연
	defer cancel()

	p := map[string]interface{}{
		"id":             testIntData,
		"DepartmentName": testStringData,
	}

	// repository의 메소드를 호출하여 DB 조회
	if Response, err := s.repo.TestSelect(ctx, p); err != nil {
		util.Logger.Println("err : ", err)
		return nil, err
	} else {
		util.Logger.Println("Response : ", Response)
		return &pb.TestResponse{
			Result: Response[0].DepartmentName,
		}, nil
	}

}
