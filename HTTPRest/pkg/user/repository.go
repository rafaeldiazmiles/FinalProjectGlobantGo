package user

import (
	"context"

	gokitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/gRPC/pkg/proto"
	"google.golang.org/grpc"
)

type grpcClient struct {
	server *grpc.ClientConn
	logger gokitlog.Logger
}

func NewgRPClient(log gokitlog.Logger, sv *grpc.ClientConn) *grpcClient {
	return &grpcClient{sv, log}
}

func (repo *grpcClient) CreateUser(ctx context.Context, rq proto.CreateUserRequest) (proto.CreateUserResponse, error) {
	logger := gokitlog.With(repo.logger, "create user", "recevied")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := proto.CreateUserRequest{
		Name:    rq.Name,
		Pwd:     rq.Pwd,
		Age:     rq.Age,
		AddInfo: rq.AddInfo,
	}

	resp, err := client.CreateUser(ctx, &protoReq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return proto.CreateUserResponse{}, err
	}

	res := proto.CreateUserResponse{
		Id: resp.Id,
	}

	return res, nil

}
