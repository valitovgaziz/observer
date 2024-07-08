package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is listening on 8282 port")
	http.HandleFunc("/GetOrderBook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/SaveOrderBook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "SavedOrderBook")
	})
	http.HandleFunc("/GetOrderHistory", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OrderHistory")
	})
	http.HandleFunc("/SaveOrder", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "SavedOrder")
	})
	http.ListenAndServe(":8282", nil)
}
