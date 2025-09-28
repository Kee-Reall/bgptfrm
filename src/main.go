package main

import (
	"blog-api/src/api"
	"blog-api/src/repository"
	"blog-api/src/utils"
	"blog-api/src/utils/logger"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	defer func() {
		if reason := recover(); reason != nil {
			if err, ok := reason.(error); ok {
				logger.ErrorException(err)
			}
			os.Exit(1)
		}
	}()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	if err := utils.ValidateEnv(); err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		panic(errors.New("Unable to connect to database: " + err.Error()))
	}

	defer func() {
		err := conn.Close(context.Background())
		if err != nil {
			logger.ErrorException(err)
		}
	}()

	repository.InitRepository(conn)

	http.Handle("/", api.InitMainRouter())

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
