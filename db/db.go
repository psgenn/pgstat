package db

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var database *sql.DB

func Open(ip string, port string, user string, password string, dbname string) {
	info := fmt.Sprintf("host=%s port=%s dbname=%s "+
	"sslmode=%s user=%s password=%s ",
		ip,
		port,
		dbname,
		"disable",
		user,
		password,
	)
	var err error
	database, err = sql.Open("postgres", info)
	if err != nil {
		log.Fatal(err)
	}
	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Get_count_connection() (uint32, error) {
	rows, err := database.Query("SELECT count(*) as total_conns FROM pg_stat_activity;")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count_connection uint32
	for rows.Next() {
		rows.Scan(&count_connection)
		fmt.Println(count_connection)
	}
	return count_connection, nil
}
