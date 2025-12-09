package constants

import (
	"database/sql"
	"fmt"
)

var DbConn *sql.DB

func InitConstants() {
	var err error
	DbConn, err = sql.Open("postgres", "")
	if err != nil {
		fmt.Printf("Error creating DB conn")
	}
}
