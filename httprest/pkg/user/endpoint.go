package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/rafaeldiazmiles/ProjectEssay/pkg/entities"
)

type CreateUserRequest struct {
	Name    string
	Pwd     string
	Age     int
	AddInfo string
}

type CreateUserResponse struct {
	Id uint32
}

type Service interface {
	CreateUser(ctx context.Context, us entities.User) (uint32, error)
}

type Endpoints struct {
	CreateUs endpoint.Endpoint
}

func MakeEndpoints(s Service) *Endpoints {

	return &Endpoints{
		CreateUs: MakeCreateUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, ok := rq.(CreateUserRequest)

		if !ok {
			return nil, errors.New("invalid request type")
		}

		res, err := s.CreateUser(ctx, entities.User{
			Name:    request.Name,
			Pwd:     request.Pwd,
			Age:     request.Age,
			AddInfo: request.AddInfo,
		})

		if err != nil {
			return nil, err
		}

		return CreateUserResponse{
			Id: res,
		}, nil

	}
}
