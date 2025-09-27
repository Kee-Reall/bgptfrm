package superAdmin

import (
	"blog-api/src/utils/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	us := make(chan int)

	go func() {
		us <- 0
	}()

	logger.Log("JAMBO NAHUY")

	if _, err := w.Write([]byte(strconv.Itoa(<-us))); err != nil {
		logger.ErrorException(err)
	}
}

func SetSuperAdminRouter(subRouter *mux.Router) {
	subRouter.Use(basicAuth)
	subRouter.Handle("/", http.HandlerFunc(getUsers)).Methods("GET")

}
