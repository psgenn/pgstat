package main

import (
	"fmt"
	//"net/http"
	//"pgstat/api"
	"github.com/psgenn/pgstat/api"
	"github.com/psgenn/pgstat/config"
	"github.com/psgenn/pgstat/db"
)

func main() {
	data, err := config.Load("./config.yaml")
	if err != nil {
		panic(err)
	}
	config, err := config.Parse(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Ip, config.Password, config.Port, config.User)

	db.Open(config.Ip, config.Port, config.User, config.Password, config.Dbname)

	db.Get_count_connection()
	api.Start()
}
