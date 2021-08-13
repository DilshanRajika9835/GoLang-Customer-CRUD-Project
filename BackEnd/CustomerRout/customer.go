package CustomerRout

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var customers []customer //Create a Customer Array

type customer struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salary  float32 `json:"salary"`
}

/*------------------------------Save Customer----------------------------------------------------*/
func SaveCustomer(w http.ResponseWriter, r *http.Request) {

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	salary := r.FormValue("salary")
	fmt.Println(salary)
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand") //connect mysql server
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?, ?, ?)", id, name, address, salary)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

/*------------------------------Delete  Customer----------------------------------------------------*/
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand") //connect mysql server
	delete, err := db.Query("delete from customer where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	json.NewEncoder(w).Encode(customers)

}

/*------------------------------Update  Customer----------------------------------------------------*/
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	salary := r.FormValue("salary")

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand") //connect mysql server
	insert, err := db.Query("UPDATE customer set name=?,address=?,salary=? where id=?", name, address, salary, id)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

/*------------------------------Get All  Customer----------------------------------------------------*/
func AllCustomer(w http.ResponseWriter, r *http.Request) {
	var allCustomer []customer
	setupCorsResponse(&w, r) //accept cross origin
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json") //set header content type application json

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand") //connect mysql server
	row, err := db.Query("Select * from customer")                     //sql get all customer
	if err != nil {
		panic(err.Error())
	} else {
		for row.Next() {
			var id string
			var name string
			var address string
			var salary float32

			err2 := row.Scan(&id, &name, &address, &salary)
			row.Columns()
			if err2 != nil {
				panic(err2.Error())
			} else {
				custList := customer{
					ID:      id,
					Name:    name,
					Address: address,
					Salary:  salary,
				}
				allCustomer = append(allCustomer, custList)
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(allCustomer)

}

/*------------------------------Search  Customer----------------------------------*/
func SearchCustomer(w http.ResponseWriter, r *http.Request) {
	var cust customer

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand")
	row, err := db.Query("select * from customer where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	} else {
		for row.Next() {
			var id string
			var name string
			var address string
			var salary float32

			err2 := row.Scan(&id, &name, &address, &salary)
			row.Columns()
			if err2 != nil {
				panic(err2.Error())
			} else {
				customer := customer{
					ID:      id,
					Name:    name,
					Address: address,
					Salary:  salary,
				}
				cust = customer
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(cust)
}

/*--------------------------------------------------------------------------------*/
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) { //allow cross origin

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //catch request method
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")

}
