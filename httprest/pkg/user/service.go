package user

import (
	"context"

	gokitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Service interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error)
}

type service struct {
	repository Service
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

func NewService(rep Service, logger gokitlog.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error) {
	logger := gokitlog.With(s.logger, "method", "CreateUser")

	userId, err := s.repository.CreateUser(ctx, CreateUserRequest{
		Name:    request.Name,
		Pwd:     request.Pwd,
		Age:     request.Age,
		AddInfo: request.AddInfo,
	})
	if err != nil {
		level.Error(logger).Log("err", err)
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{Id: userId.Id}, nil
}
