package repository

import (
	"blog-api/src/shared"
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	getUsersQuery         = "SELECT email, login, created_at FROM users where deleted_at is null order by created_at DESC;"
	getUserWithLoginQuery = "SELECT email, login, created_at, deleted_at from users where login = @login;"
)

func (r *Repository) GetUsers(pagParams *shared.PaginationParams) ([]shared.User, error) {
	rows, err := r.conn.Query(context.Background(), getUsersQuery)
	if err != nil {
		return nil, connectionError
	}
	defer rows.Close()

	userCap := pagParams.PageSize

	if userCap > 128 {
		userCap = 128
	}

	users := make([]shared.User, 0, userCap)
	for rows.Next() {
		var u shared.User
		if err := rows.Scan(&u.Email, &u.Login, &u.CreatedAt); err != nil {
			return nil, connectionError
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repository) GetUserByLogin(login string) (*shared.User, error) {
	rows, err := r.conn.Query(context.Background(), getUserWithLoginQuery, pgx.NamedArgs{"login": login})
	if err != nil {
		return nil, connectionError
	}

	if !rows.Next() {
		return nil, notFound
	}

	var u shared.User
	if err := rows.Scan(&u.Email, &u.Login, u.CreatedAt, u.DeletedAt); err != nil {
		return nil, connectionError
	}
	return &u, nil
}
