package handler

import (
	"fmt"
	"net/http"
)
// define the order type :


type Order struct {}


// define handlers for crud methods :

// Create handler :
func (order *Order) Create( w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new Order ")
} 

// Get handler :
func (order *Order) List( w http.ResponseWriter, r *http.Request) {
	fmt.Println("List  all orders ")
}


// Get by id handler :
func (order *Order) GetById( w http.ResponseWriter, r *http.Request) {
	fmt.Println("get order by id ")
}

// Update by id handler :
func (order *Order) UpdateById( w http.ResponseWriter, r *http.Request) {
	fmt.Println("update order by id ")
}

// Delete by id handler :
func (order *Order) DeleteById( w http.ResponseWriter, r *http.Request) {
	fmt.Println("update order by id ")
}





