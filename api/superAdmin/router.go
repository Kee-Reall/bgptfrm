package superAdmin

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUsers(w http.ResponseWriter, r *http.Request)  {
	us := make(chan int)

	go func() {
		us <- 0
	}()

	w.Write([]byte(strconv.Itoa(<-us)))
}

func SetSuperAdminRouter(subRouter *mux.Router) {
	subRouter.Methods("GET").HandlerFunc(getUsers)
}
