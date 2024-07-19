package service

import (
	"context"

	pb "ultraman/api/consume/v1"
)

type ConsumeService struct {
	pb.UnimplementedConsumeServer
}

func NewConsumeService() *ConsumeService {

	return &ConsumeService{}
}

func (s *ConsumeService) GetConsume(ctx context.Context, req *pb.ConsumeRequest) (*pb.GetConsumeReply, error) {
	return &pb.GetConsumeReply{}, nil
}
func (s *ConsumeService) GetMetrics(ctx context.Context, req *pb.MetricsRequest) (*pb.MetricsReply, error) {
	return &pb.MetricsReply{}, nil
}
