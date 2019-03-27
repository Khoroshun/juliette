package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khoroshun/juliette/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	//router.HandleFunc("/api/user/{id:[0-9]+}/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.HandleFunc("/api/order/new", controllers.CreateOrder).Methods("POST")


	router.HandleFunc("/api/bonusaccount/get", controllers.GetBonusAccount).Methods("GET")
	//router.HandleFunc("/api/bonustransaction/new", controllers.CreateBonusTransaction).Methods("POST")
	router.HandleFunc("/api/bonustransaction/get", controllers.GetBonusTransaction).Methods("GET")


	router.HandleFunc("/api/discountaccount/get", controllers.GetDiscountAccount).Methods("GET")
	router.HandleFunc("/api/discountchanges/new", controllers.CreateDiscountChanges).Methods("POST")
	router.HandleFunc("/api/discountchanges/get", controllers.GetDiscountChanges).Methods("GET")

	//router.Use(app.JwtAuthentication) //attach JWT auth middleware

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