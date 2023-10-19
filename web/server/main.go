package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("The Server was Pinged!")
	// res, err := httputil.DumpRequest(r, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(string(res))

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}
	fmt.Println(r.PostForm)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
	mux.HandleFunc("/plm/cors", Cors)
	mux.HandleFunc("/data", handleRequest)
	http.ListenAndServe(":8080", mux)
}
