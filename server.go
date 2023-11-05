package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

// func Cors(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8; charset=ascii")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
// 	w.Write([]byte("Hello, World!"))
// }

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET here")
		tpl.ExecuteTemplate(w, "login.html", nil)
	} else if r.Method == "POST" {
		fmt.Println("POST here")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form")
		}
		fmt.Println(r.PostForm["username"][0])
		fmt.Println(r.PostForm["passcode"][0])

		http.Redirect(w, r, "/index", http.StatusSeeOther)
		w.WriteHeader(http.StatusSeeOther)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		fmt.Println("Redirect successful")
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		fmt.Println("Error parsing form")
// 	}
// 	fmt.Println(r.PostForm["username"][0])
// 	fmt.Println(r.PostForm["passcode"][0])

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	resp := make(map[string]string)
// 	resp["message"] = "Status OK"
// 	jsonResp, err := json.Marshal(resp)
// 	if err != nil {
// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 	}

// 	w.Write(jsonResp)
// 	return
// }

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// mux.HandleFunc("/cors", Cors)
	mux.HandleFunc("/index", indexPage)
	mux.HandleFunc("/login", loginPage)
	http.ListenAndServe(":9000", mux)
}
