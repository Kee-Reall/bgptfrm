package repository

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	Repo            *Repository
	connectionError = errors.New("connection error")
	notFound        = errors.New("required data not found")
)

type Repository struct {
	conn *pgx.Conn
}

func InitRepository(conn *pgx.Conn) bool {
	if Repo == nil {
		if conn == nil {
			return false
		}
		Repo = &Repository{conn}
	}
	return true
}

func GetConnection() *pgx.Conn {
	return Repo.conn
}
