package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/pkg/entities" // "github.com/junereycasuga/gokit-grpc-demo/service"
)

// Service interface describes actions on Users
// Users(Store) - defines the interface we expect our database implementation to follow
type Service interface {
	// Authenticate(id string) error
	CreateUser(context.Context, entities.User) (uint32, error)
	// UpdateUser(usr User) error
	// GetUser(id int32) (User, error)
	// DeleteUser(id int32) error
}

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	CreateUser endpoint.Endpoint
}

// CreateUserRequest struct holds the endpoint request definition
type CreateUserRequest struct {
	pwd     string
	name    string
	age     uint32
	addInfo string
	// parents []Parent   --> Para implementar cuando haya parents
}

// CreateUserResponse struct holds the endpoint response definition
type CreateUserResponse struct {
	id uint32
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateUserRequest)
		if !ok {
			return nil, err //Tengo que retornar un error custom
		}
		result, err := s.CreateUser(ctx, entities.User{
			Pwd:     req.pwd,
			Name:    req.name,
			Age:     req.age,
			AddInfo: req.addInfo,
		})
		if err != nil {
			return nil, err
		}
		return CreateUserResponse{id: result}, nil
	}
}
