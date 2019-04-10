package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khoroshun/juliette/app"
	"github.com/khoroshun/juliette/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	// Client
	router.HandleFunc("/api/client/get", controllers.GetClientHandler).Methods("GET")
	router.HandleFunc("/api/client/new", controllers.CreateClientHandler).Methods("POST")
	router.HandleFunc("/api/client/{id:[0-9]+}/update", controllers.UpdateClientHandler).Methods("POST")

	// Order
	router.HandleFunc("/api/order/get", controllers.GetOrderHandler).Methods("GET")
	router.HandleFunc("/api/order/new", controllers.CreateOrderHandler).Methods("POST")
	router.HandleFunc("/api/order/{order_num}/update", controllers.UpdateOrderHandler).Methods("POST")

	// Bonus Transaction
	router.HandleFunc("/api/order/get", controllers.GetBonusTransactionsHandler).Methods("GET")
	router.HandleFunc("/api/order/new", controllers.CreateBonusTransactionHandler).Methods("POST")
	router.HandleFunc("/api/order/update", controllers.UpdateBonusTransactionHandler).Methods("POST")

/*
	router.HandleFunc("/api/order/new", controllers.CreateOrder).Methods("POST")

	router.HandleFunc("/api/bonustransaction/new", controllers.CreateBonusTransaction).Methods("POST")
	router.HandleFunc("/api/bonusaccount/get", controllers.GetBonusAccount).Methods("GET")
	router.HandleFunc("/api/bonustransaction/get", controllers.GetBonusTransactions).Methods("GET")


	router.HandleFunc("/api/discountchanges/new", controllers.CreateDiscountChanges).Methods("POST")
	router.HandleFunc("/api/discountaccount/get", controllers.GetDiscountAccount).Methods("GET")
	router.HandleFunc("/api/discountchanges/get", controllers.GetDiscountChanges).Methods("GET")
*/
	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}
	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}