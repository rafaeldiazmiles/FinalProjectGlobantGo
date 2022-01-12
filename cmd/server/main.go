package main

import (
	"database/sql"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rafaeldiazmiles/ProjectEssay/pkg/user"
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
	defer level.Info(logger).Log("msg", "Service ended")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.01:5050)/projectDB")
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	defer db.Close()

	{
		repo := user.NewRepo(db, logger)
		srv := user.NewService(repo, logger)
		endP := user.MakeEndpoints(srv)
	}
}
