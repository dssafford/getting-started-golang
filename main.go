package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'TOAST','JJJJJJJJJJHEY!!!!!','DUDEEEE','San Francisco, Amsterdam, Berlin, New York','Tokyo' 'Eat shit', 'fuck you'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
