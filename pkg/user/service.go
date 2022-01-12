package user

import (
	"context"

	"github.com/go-kit/log"
	"github.com/rafaeldiazmiles/ProjectEssay/pkg/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, us entities.User) (int32, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService returns a Service with all of the expected dependencies
func NewService(rep Repository, logger log.Logger) *service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, ent entities.User) (int32, error) {
	usrID, err := s.repository.CreateUser(ctx, ent)
	if err != nil {
		return -1, err
	}
	return usrID, nil
}
