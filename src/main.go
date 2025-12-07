package main

import (
	"blog-api/db"
	"blog-api/src/api"
	"blog-api/src/core"
	"blog-api/src/repository"
	"blog-api/src/utils/logger"
	"blog-api/src/utils/validation"
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	defer func() {
		if reason := recover(); reason != nil {
			log.Fatal(reason)
		}
	}()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	if err := validation.ValidateEnv(); err != nil {
		panic(err)
	}

	ctx, dbUrl := context.Background(), os.Getenv("DB_URL")

	if err := db.StartMigration(dbUrl); err != nil {
		panic("migration failed" + err.Error())
	}

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		panic(errors.New("Unable to connect to database: " + err.Error()))
	}

	defer func() {
		err := conn.Close(ctx)
		if err != nil {
			logger.ErrorException(err)
		}
	}()

	if err := conn.Ping(ctx); err != nil {
		panic(errors.New("Control ping was failure with: " + err.Error()))
	}

	if ok := repository.InitRepository(conn); !ok {
		panic("can't initialize db connection")
	}

	if ok := core.Init(); !ok {
		panic("can't initialize core module")
	}

	http.Handle("/", api.InitMainRouter())

	if err := http.ListenAndServe(":"+os.Getenv("APP_ADDR"), nil); err != nil {
		panic(err)
	}
}
