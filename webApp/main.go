// package main import ( "fmt" "log" "net/http" ) const ( // Host name of the HTTP Server Host = "localhost" // Port of the HTTP Server Port = "8080" ) func home(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "HOME Page") } func about(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "ABOUT Page") } func services(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "SERVICES Page") } func main() { http.HandleFunc("/", home) http.HandleFunc("/about", about) http.HandleFunc("/services", services) err := http.ListenAndServe(Host+":"+Port, nil) if err != nil { log.Fatal("Error Starting the HTTP Server : ", err) return } }
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json: "email"`
	Phone string `json:"phone"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "me@lol.net",
		Phone: "12345",
	}
	json.NewEncoder(w).Encode(user)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(w, "Failed to load template")
	}

	user := User{
		Id:    1,
		Name:  "John",
		Email: "me@lol.net",
		Phone: "1234",
	}

	t.Execute(w, user)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT Page")
}

func services(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SERVICES Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/services", services)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/template", templateHandler)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
