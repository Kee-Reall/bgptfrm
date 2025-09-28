package superAdmin

import (
	"blog-api/src/repository"
	"context"
	"encoding/json"
	"net/http"
)

type User struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	conn, err := repository.GetConnection()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	rows, err := conn.Query(context.Background(), "SELECT id, title, name FROM users")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer rows.Close()

	users := make([]User, 0, 5)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Title, &u.Name)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		users = append(users, u)
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
