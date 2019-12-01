package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khoroshun/juliette/app"
	"github.com/khoroshun/juliette/controllers"
	cron "github.com/robfig/cron"
	sms "github.com/wildsurfer/turbosms-go"
	"net/http"
	"os"
)

func main() {


	c := cron.New()
	c.AddFunc("* */1 * * *", func() {





		//sender := sms.NewClient("JulietteBrand", "0997740160jb")
		//sender.SendSMS("Juliette", "+380967154107", "test", "")

	})
	c.Start()






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
	router.HandleFunc("/api/bonustransaction/get", controllers.GetBonusTransactionsHandler).Methods("GET")
	router.HandleFunc("/api/bonustransaction/new", controllers.CreateBonusTransactionHandler).Methods("POST")
	router.HandleFunc("/api/bonustransaction/update", controllers.UpdateBonusTransactionHandler).Methods("POST")


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