package user

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/gRPC/pkg/proto"
)

type gRPCServer struct {
	createUser gt.Handler
	proto.UnimplementedUserServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints Endpoints, logger log.Logger) proto.UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.CreateUserResponse), nil
}

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.CreateUserRequest)
	return CreateUserRequest{
		pwd:     req.Pwd,
		name:    req.Name,
		age:     req.Age,
		addInfo: req.AddInfo}, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateUserResponse)
	return &proto.CreateUserResponse{
		Id:     resp.id,
		Status: &proto.Status{},
	}, nil
}
