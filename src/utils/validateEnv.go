package utils

import (
	"errors"
	"net/url"
	"os"
	"strconv"
)

func ValidateEnv() error {

	for _, checker := range []func() bool{checkAdminCred, checkPort, checkDbUrl} {
		if ok := checker(); !ok {
			return errors.New("at least one of env cred is not defined or has not allowed value")
		}
	}
	return nil
}

func checkAdminCred() bool {
	var minLength int = 3
	adminLogin, adminPassword := os.Getenv("ADMIN_LOGIN"), os.Getenv("ADMIN_PASSWORD")
	return len([]rune(adminLogin)) > minLength || len([]rune(adminPassword)) > minLength
}

func checkPort() bool {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return false
	}
	portNum, err := strconv.Atoi(portStr)
	if err != nil {
		return false
	}

	return portNum > 0 && portNum <= 65535
}

func checkDbUrl() bool {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return false
	}
	uri, err := url.Parse(dbUrl)
	if err != nil || uri == nil {
		return false
	}

	if uri.Scheme != "postgres" {
		return false
	}
	return true

}
