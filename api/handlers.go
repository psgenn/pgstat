package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/psgenn/pgstat/db"
	"fmt"
	"os"
)

type Response struct {
	//{"status": "ok", "response": {"name_response": <value_respose>}}
	Status string        `json:"status"`
	Response interface{} `json:"response"`
}

type Error struct {
	//{"status": "error", "msg": <message>}
	Status string `json:"status"`
	Msg string    `json:"msg"`
}

func make_response(response interface{}) Response {
	var res Response
	res.Status = "ok"
	res.Response = response
	return res
}

func make_error(msg string) Error {
	var err Error
	err.Status = "error"
	err.Msg = msg
	return err
}

func count_connection(w http.ResponseWriter, r *http.Request) {
	var str string
	if len(r.URL.RawQuery) > 0 {
		str = r.URL.Query().Get("name")
		if str == "" {
			w.WriteHeader(400)
			return
		}
	}
	count, err := db.Get_count_connection()
	var res interface{}
	if err != nil {
		res = make_error("not connected")
		fmt.Println("not connected")
	} else {
		response := make(map[string]uint32)
		response["count_connection"] = count
		res = make_response(response)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(500)
	}
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/count_connection", count_connection).Methods("GET")
	http.ListenAndServe(":8080", router)
}
