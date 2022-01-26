package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/HTTPRest/pkg/entities"
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
	CreateUser endpoint.Endpoint
}

func MakeEndpoints(s Service) *Endpoints {

	return &Endpoints{
		CreateUser: MakeCreateUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, ok := rq.(CreateUserRequest)

		if !ok {
			return nil, errors.New("invalid request type")
		}

		res, err := s.CreateUser(ctx, request)

		if err != nil {
			return nil, err
		}

		return CreateUserResponse{
			Id: res.Id,
		}, nil

	}
}
