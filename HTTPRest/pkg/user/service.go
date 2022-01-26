package user

import (
	"context"

	gokitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/gRPC/pkg/proto"
)

type Service interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error)
}

type service struct {
	repository *grpcClient
	logger     gokitlog.Logger
}

type CreateUserRequest struct {
	Name    string
	Pwd     string
	Age     int
	AddInfo string
}

type CreateUserResponse struct {
	Id uint32
}

func NewService(rep *grpcClient, logger gokitlog.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error) {
	logger := gokitlog.With(s.logger, "method", "CreateUser")

	userId, err := s.repository.CreateUser(ctx, proto.CreateUserRequest{
		Name:    request.Name,
		Pwd:     request.Pwd,
		Age:     uint32(request.Age),
		AddInfo: request.AddInfo,
	})
	if err != nil {
		level.Error(logger).Log("err", err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{Id: userId.Id}, nil
}
