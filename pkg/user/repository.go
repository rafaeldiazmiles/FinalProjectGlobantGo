//go:generate mockgen -destination=repositories_mocks_test.go -package=rocket https://github.com/rafaeldiazmiles/FinalProjectGlobantGo/ Users
package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/pkg/entities"
)

var CreateUserQuery string = "INSERT INTO USER (name, pwd, age, add_info) VALUES (?,?,?,?,?)"

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
	res, err := stmt.ExecContext(ctx, entities.User.Name, entities.User.Pwd, entities.User.Age, entities.User.AddInfo)
	if err != nil {
		level.Error(repo.Logger).Log(err)
		return 0, err
	}

	return 0, nil

}

// func (repo *sqlRepo) CreateUser(ctx context.Context, user entities.User, newId string) (string, error) {

// 	repo.Logger.Log(repo.Logger, "Repository method", "Create user")

// 	stmt, err := repo.DB.PrepareContext(ctx, utils.CreateUserQuery)
// 	if err != nil {
// 		level.Error(repo.Logger).Log(err)
// 		return "", err
// 	}

// 	defer stmt.Close()
// 	res, err := stmt.ExecContext(ctx, user.Name, newId, user.Pass, user.Age, user.Email)
// 	if err != nil {
// 		level.Error(repo.Logger).Log(err)
// 		return "", err
// 	}

// 	repo.Logger.Log(repo.Logger, res, "rows affected")

// 	return newId, nil
// }
