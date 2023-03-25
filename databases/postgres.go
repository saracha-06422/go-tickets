package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dew12345"
	dbname   = "go_booking_tickets"
)

func ConnectPostgres() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully database connected!")
}
