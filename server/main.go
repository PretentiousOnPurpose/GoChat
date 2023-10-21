package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}
	fmt.Println(r.PostForm["username"][0])
	fmt.Println(r.PostForm["passcode"][0])

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cors", Cors)
	mux.HandleFunc("/data", handleRequest)
	mux.HandleFunc("/login", handleRequest)
	http.ListenAndServe(":8080", mux)
}
