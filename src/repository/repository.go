package repository

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var currentRepository *Repository

type Repository struct {
	conn *pgx.Conn
}

func (r *Repository) getConn() *pgx.Conn {
	return r.conn
}

func InitRepository(conn *pgx.Conn) bool {
	if currentRepository == nil {
		if conn == nil {
			return false
		}
		currentRepository = &Repository{conn}
		return true
	}
	return true
}

func GetConnection() (*pgx.Conn, error) {
	if currentRepository == nil {
		return nil, errors.New("connection is not initialized")
	}
	return currentRepository.getConn(), nil
}
