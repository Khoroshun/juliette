package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khoroshun/juliette/app"
	"github.com/khoroshun/juliette/controllers"
	"github.com/khoroshun/juliette/models"
	"github.com/robfig/cron"
	sms "github.com/wildsurfer/turbosms-go"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func  getPhoneByBonusAccountID(id uint) string {

	bonusAccount := &models.BonusAccount{}
	errAc := models.GetDB().Table("bonus_accounts").Where("id = ?", id).Find(bonusAccount).Error
	if errAc != nil {
		fmt.Println(errAc)
	}

	client := &models.Client{}
	errCl := models.GetDB().Table("clients").Where("id = ?",bonusAccount.Client).Find(client).Error
	if errCl != nil {
		fmt.Println(errCl)
	}

	return client.Phone
}

func sendPushSmsAccount(days int64, message string)  {


	t := time.Now()
	start := t.Unix() - (86400*days)
	t_start := time.Unix(start, 0)
	u_start := t_start.Format("2006-01-02 15:04:05")

	end := t.Unix() - (86400*(days+1))
	t_end := time.Unix(end, 0)
	u_end := t_end.Format("2006-01-02 15:04:05")

	finish := t.Unix() + (86400*(180 - days))
	t_finish := time.Unix(finish, 0)
	u_finish := t_finish.Format("02.01.2006")

	bonusAccounts := make([]*models.BonusAccount, 0)
	err := models.GetDB().Table("bonus_accounts").Where("updated_at < ? AND updated_at > ?", u_start, u_end).Find(&bonusAccounts).Error
	if err != nil {
		fmt.Println(err)
	}

	message = strings.Replace(message, "#FINISH#", u_finish, -1)

	sender := sms.NewClient("JulietteBrand", "0997740160jb")
	for i := 0; i < len(bonusAccounts); i++ {

		message = strings.Replace(message, "#SUMM#", strconv.Itoa(bonusAccounts[i].Summ), -1)
		fmt.Println(getPhoneByBonusAccountID(bonusAccounts[i].ID))
		fmt.Println(message)

		sender.SendSMS("Juliette", getPhoneByBonusAccountID(bonusAccounts[i].ID), message, "")
	}


}

func bonusActivate(){

	var days int64
	message := "Ваши бонусы активированы - https://juliette-sun.com.ua/check_bonus.php"
	days = 10
	t := time.Now()
	start := t.Unix() - (86400*days)
	t_start := time.Unix(start, 0)
	u_start := t_start.Format("2006-01-02 15:04:05")

	end := t.Unix() - (86400*(days+1))
	t_end := time.Unix(end, 0)
	u_end := t_end.Format("2006-01-02 15:04:05")

	bonusTransactions := make([]*models.BonusTransaction, 0)
	err := models.GetDB().Table("bonus_transactions").Where("created_at < ? AND created_at > ?", u_start, u_end).Find(&bonusTransactions).Error
	if err != nil {
		fmt.Println(err)
	}

	sender := sms.NewClient("JulietteBrand", "0997740160jb")
	for i := 0; i < len(bonusTransactions); i++ {

		if !bonusTransactions[i].Active {

			bonusTransactions[i].Active = true
			bonusTransactions[i].Update()

			bonusAccount := &models.BonusAccount{}
			errAc := models.GetDB().Table("bonus_accounts").Where("id = ?", bonusTransactions[i].Account).Find(bonusAccount).Error
			if errAc != nil {
				fmt.Println(errAc)
			}

			sum := bonusAccount.Summ + bonusTransactions[i].Summ
			models.GetDB().Table("bonus_accounts").Where("id = ?", bonusAccount.ID).Update("summ", sum)

			fmt.Println(getPhoneByBonusAccountID(bonusTransactions[i].Account))
			fmt.Println(message)
			sender.SendSMS("Juliette", getPhoneByBonusAccountID(bonusTransactions[i].Account), message, "")
		}

	}


}

func main() {



	c := cron.New()
	c.AddFunc("0 0 * * *", func() { // раз в сутки

		bonusActivate()
		sendPushSmsAccount(166,"Срок действия Ваших бонусов истекает через 14 дней. На Вашем счете #SUMM# грн") // 14 дней до конца
		sendPushSmsAccount(173,"Покупай больше- плати меньше. Приглашаем вас воспользоваться вашими бонусами на приятные покупки. Доступно #SUMM# грн. #FINISH# бонусы будут деактивированы") // 3 дней до конца

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

	err_serv := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err_serv != nil {
		fmt.Print(err_serv)
	}
}