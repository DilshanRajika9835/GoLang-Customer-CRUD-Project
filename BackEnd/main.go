package main

import (
	"database/sql"
	"encoding/json"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
var customers []customer

type customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Salary float32 `json:"salary"`

}

func saveCustomer(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	name :=r.FormValue("name")
	address :=r.FormValue("address")
	salary:=r.FormValue("salary")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/test")
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?, ?, ?)",id,name, address,salary)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/test")
	delete, err := db.Query("delete from customer where id=?",id)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(delete)
	defer delete.Close()
	json.NewEncoder(w).Encode(customers)

}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	name :=r.FormValue("name")
	address :=r.FormValue("address")
	salary:=r.FormValue("salary")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/test")
	insert, err := db.Query("UPDATE customer set name=?,address=?,salary=? where id=?",name, address,salary,id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
func allCustomer(w http.ResponseWriter, r *http.Request) {
	var	allCustomer []customer
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/test")
	row, err := db.Query("Select * from customer")
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var id string
			var name string
			var address string
			var salary float32

			err2 := row.Scan(&id, &name, &address, &salary)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				custList := customer{
					ID:id,
					Name:name,
					Address: address,
					Salary: salary,
				}
				allCustomer = append(allCustomer, custList)
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(allCustomer)



}
func main() {

	r := mux.NewRouter()
	fmt.Println("Server Running...")
	r.HandleFunc("/api/customer", saveCustomer).Methods("POST","OPTIONS")
	r.HandleFunc("/api/customer", deleteCustomer).Methods("DELETE","OPTIONS")
	r.HandleFunc("/api/customer", updateCustomer).Methods("PUT","OPTIONS")
	r.HandleFunc("/api/customer", allCustomer).Methods("GET","OPTIONS")
	log.Fatal(http.ListenAndServe(":8000", r))


}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")


}
