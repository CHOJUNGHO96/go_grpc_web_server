package usecase

import (
	"context"
	pb "sequence_game_server/api/v1/test2"
	"sequence_game_server/pkg/test2/model"
	repository "sequence_game_server/pkg/test2/repository"
	"time"
)

type Response struct {
	Test2Response []model.SelectCompanyDepartment
}

type Test2Service struct {
	pb.UnimplementedTest2ServiceServer
	repo *repository.Repository
}

func NewTestService(repo *repository.Repository) *Test2Service {
	return &Test2Service{repo: repo}
}

func (s *Test2Service) Test2(ctx context.Context, req *pb.Test2Request) (*pb.Test2Response, error) {
	test2IntData := req.TestIntData

	// 5초 제한의 새로운 컨텍스트 생성
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// 리소스 반환을 위한 cancel 함수 호출 지연
	defer cancel()

	p := map[string]interface{}{
		"id": test2IntData,
	}

	// repository의 메소드를 호출하여 DB 조회
	if Response, err := s.repo.Test2Select(ctx, p); err != nil {
		return nil, err
	} else {
		return &pb.Test2Response{
			TestMapData: map[string]string{
				"result": Response[0].DepartmentName,
			},
		}, nil
	}

}
