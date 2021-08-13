package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"./CustomerRout"
)

func main() {

	r := mux.NewRouter()
	fmt.Println("Server Running...")
	r.HandleFunc("/api/customer", CustomerRout.SaveCustomer).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.DeleteCustomer).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.UpdateCustomer).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/customer", CustomerRout.AllCustomer).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/customer/{id}", CustomerRout.SearchCustomer).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", r))

}
