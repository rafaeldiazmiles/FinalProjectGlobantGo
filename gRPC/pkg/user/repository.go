package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/gRPC/pkg/entities"
)

type SQLRepo struct {
	DB     *sql.DB
	Logger log.Logger
}

func NewRepo(sql *sql.DB, logger log.Logger) *SQLRepo {
	return &SQLRepo{
		DB:     sql,
		Logger: log.With(logger, "error", "db"),
	}
}

func (repo *SQLRepo) CreateUser(ctx context.Context, us entities.User) (uint32, error) {

	repo.Logger.Log(repo.Logger, "Repository method", "Create user")
	stmt, err := repo.DB.PrepareContext(ctx, CreateUserQuery)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, us.Name, us.Pwd, us.Age, us.AddInfo)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return 0, err
	}

	repo.Logger.Log(repo.Logger, res, "rows affected")
	newID, _ := res.LastInsertId()
	return uint32(newID), nil

}
