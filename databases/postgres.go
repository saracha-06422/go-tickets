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
	password = "0612785129"
	dbname   = "Tickets"
)

func ConnectPostgres() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully database connected!")

	rows, err := db.Query("select * from public.users order by id ASC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var userName string
		var email string
		var tel string

		err := rows.Scan(&id, &userName, &email, &tel)
		if err != nil {
			panic(err)
		}

		fmt.Printf("id: %d, userName: %s, email: %s, tel: %s\n", id, userName, email, tel)
	}
}
