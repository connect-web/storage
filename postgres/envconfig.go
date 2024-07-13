package postgres

import (
	"fmt"
	"os"
)

func GetUrl() string {
	envVar := os.Getenv("siteonline")

	var user, password, host, port, dbname string

	if envVar == "site" {
		user = os.Getenv("lowLatencyUser")
		password = os.Getenv("lowLatencyPassword")
		host = os.Getenv("lowLatencyHost")
		port = os.Getenv("lowLatencyPort")
		dbname = os.Getenv("lowLatencyDatabase")
	} else {
		// LOCALHOST connect to server credentials
		user = os.Getenv("lowLatencyWebUser")
		password = os.Getenv("lowLatencyWebPassword")
		host = os.Getenv("lowLatencyWebHost")
		port = os.Getenv("lowLatencyWebPort")
		dbname = os.Getenv("lowLatencyWebDatabase")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
