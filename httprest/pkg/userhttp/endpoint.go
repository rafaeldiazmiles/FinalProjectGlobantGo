package userhttp

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/pkg/entities"
)

type HTTPCreateUserRequest struct {
	Name    string
	Pwd     string
	Age     int
	AddInfo string
}

type HTTPCreateUserResponse struct {
	Id uint32
}

type HTTPService interface {
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
