package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Konstanta100/home_work_basic/hw13_http/entity"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host")
	port := flag.String("port", "8888", "port")
	flag.Parse()

	fmt.Println("Server start")
	http.HandleFunc("/", middleware)
	http.HandleFunc("/about", about)
	http.HandleFunc("/user/create", createUser)
	http.ListenAndServe(*host+":"+*port, nil) //nolint:gosec
}

func middleware(w http.ResponseWriter, r *http.Request) {
	printInfoRequest(r)

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Allowed only methods: GET and POST", http.StatusMethodNotAllowed)
		return
	}
}

func printInfoRequest(r *http.Request) {
	fmt.Println("[INFO request]", r.Host, r.URL.Path, r.Method)
}

func about(w http.ResponseWriter, r *http.Request) {
	printInfoRequest(r)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	page := entity.Page{
		Name:     "About",
		Host:     r.Host,
		Resource: r.URL.Path,
		Method:   r.Method,
	}

	json.NewEncoder(w).Encode(page)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	printInfoRequest(r)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser entity.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	fmt.Printf("New user: %+v\n", newUser)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newUser)
}
