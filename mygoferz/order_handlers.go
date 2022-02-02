package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Order struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

var orders []Order

func init() {
	loadDbJson()
}

func saveDbJson() {
	file, _ := json.MarshalIndent(orders, "", " ")
	_ = ioutil.WriteFile("./orders.json", file, 0644)
}

func loadDbJson() {
	file, err := ioutil.ReadFile("./orders.json")
	if err != nil {
		fmt.Println("First time run, initialising orders db")
		return
	}
	_ = json.Unmarshal([]byte(file), &orders)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderListBytes, err := json.Marshal(orders)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(orderListBytes)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := Order{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order.Name = r.Form.Get("name")
	order.Description = r.Form.Get("description")
	order.Date = r.Form.Get("date")

	orders = append(orders, order)
	saveDbJson()

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"Name": order.Name,
		"Desc": order.Description,
		"Date": order.Date,
	}).Info("Create order")
	//persist to disk
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
