//go:generate mockgen -destination=repositories_mocks_test.go -package=rocket https://github.com/rafaeldiazmiles/ProjectEssay Users
package user

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	"github.com/rafaeldiazmiles/ProjectEssay/pkg/entities"
)

type SQLRepo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(sql *sql.DB, logger log.Logger) *SQLRepo {
	return &SQLRepo{
		db:     sql,
		logger: log.With(logger, "error", "db"),
	}
}

func (repo *SQLRepo) CreateUser(ctx context.Context, us entities.User) (uint32, error) {
	return 0, nil

}

// // NewDatabase - returns a pointer to a database object
// func NewDatabase() (*gorm.DB, error) {
// 	fmt.Println("Setting up new database connection")

// 	dbUsername := "rafaeldiaz" //os.Getenv("DB_USERNAME")
// 	dbPassword := "postgres" //os.Getenv("DB_PASSWORD")
// 	dbHost := "172.17.0.2"   //os.Getenv("DB_HOST") //172.17.0.2
// 	dbTable := "comments"    //os.Getenv("DB_TABLE")
// 	dbPort := "5432"         //os.Getenv("DB_PORT") //5432

// 	// connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

// 	connectString := fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=disable", dbHost, dbPort, dbTable, dbPassword, dbUsername)

// 	db, err := gorm.Open("postgres", connectString)
// 	if err != nil {
// 		return db, err
// 	}

// 	if err := db.DB().Ping(); err != nil {
// 		return db, err
// 	}

// 	return db, nil
// }

// // // Service - our Users service, responsible for updating the Users DB
// // type Service struct {
// // 	User Users
// // }

// // New - returns a new instance of our Users service
// // func New(user Users) Service {
// // 	return Service{
// // 		User: user,
// // 	}
// // }

// // // GetUser - retrieves a user based on the ID from the
// // func (s Service) GetUser(ctx context.Context, id int32) (User, error) {
// // 	usr, err := s.User.GetUser(id)
// // 	if err != nil {
// // 		return User{}, err
// // 	}
// // 	return usr, nil
// // }

// // // DeleteRocket - deletes a rocket from our inventory
// // func (s Service) DeleteRocket(ctx context.Context, id int32) error {
// // 	err := s.User.DeleteUser(id)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }
