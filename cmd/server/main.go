package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	grpc "google.golang.org/grpc"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/pkg/user"
	"github.com/rafaeldiazmiles/FinalProjectGlobantGo/proto"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "Service started")
	// defer level.Info(logger).Log("msg", "Service ended")

	db, err := sql.Open("mysql", "root:Password1*@tcp(127.0.01:3306)/users")
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	defer db.Close()
	fmt.Print("Estoy aca")
	repo := user.NewRepo(db, logger)
	srv := user.NewService(repo, logger)
	endP := user.MakeEndpoints(srv)
	trnsp := user.NewGRPCServer(endP, logger)

	fmt.Print("Estoy aca en segundo lugar")

	listener, err := net.Listen("tcp", ":50005")
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		os.Exit(-1)

	}
	fmt.Print("Estoy aca en tercer lugar")

	// go func() {

	baseServer := grpc.NewServer()
	proto.RegisterUserServiceServer(baseServer, trnsp)
	fmt.Print("Estoy aca en ultimo lugar")
	baseServer.Serve(listener)
	// }()

}
