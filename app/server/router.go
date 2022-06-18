package server

import (
	"log"
	"net/http"
	"os"
	"ransmart_pay/app/handler/loginHandler"
	"ransmart_pay/app/handler/payHandler"
	"ransmart_pay/app/handler/payHistoryHandler"
	"ransmart_pay/app/handler/tokenHandler"
	"ransmart_pay/app/handler/userHandler"
	"ransmart_pay/app/helper/helper"
	"ransmart_pay/app/helper/response"
	"ransmart_pay/app/models/payHistoryModel"
	"ransmart_pay/app/models/payModel"
	"ransmart_pay/app/models/userModel"
	"ransmart_pay/app/repository"
	"ransmart_pay/app/repository/payHistoryRepository"
	"ransmart_pay/app/repository/payRepository"
	"ransmart_pay/app/repository/userRepository"
	"ransmart_pay/app/service"
	"ransmart_pay/app/service/payHistoryService"
	"ransmart_pay/app/service/payService"
	"ransmart_pay/app/service/userService"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {
	// try connect to database
	log.Println("Connecting to Database...")
	db, err := gorm.Open(mysql.Open(getConnectionString()), &gorm.Config{})
	helper.CheckFatal(err)

	// migrate model to database
	db.AutoMigrate(&userModel.User{}, &payModel.PayModel{}, &payHistoryModel.PayHistoryModel{})
	log.Println("Database Connected")

	// generate repository
	allRepositories := repository.Repository{
		IUserRepository:       userRepository.NewRepository(db),
		IPayRepository:        payRepository.NewRepository(db),
		IPayHistoryRepository: payHistoryRepository.NewRepository(db),
	}

	// try connect to redis
	// log.Println("Connecting to Redis in Background...")
	// redis := connectToRedis()

	// generate service
	allServices := service.Service{
		IUserService:       userService.NewService(allRepositories, db),
		IPayService:        payService.NewService(allRepositories),
		IPayHistoryService: payHistoryService.NewService(allRepositories, db),
	}

	// generate handler
	user := userHandler.NewUserHandler(allServices)
	login := loginHandler.NewLoginHandler(allServices)
	pay := payHandler.NewPayHandler(allServices)
	payHistory := payHistoryHandler.NewPayHistoryHandler(allServices)

	// router
	r := chi.NewRouter()

	// check service
	r.Group(func(g chi.Router) {
		g.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response.ResponseRunningService(w)
		})
	})

	// // global token
	// r.Group(func(g chi.Router) {
	// 	g.Get("/globaltoken", login.GenerateToken)
	// })

	// login
	r.Group(func(l chi.Router) {
		l.Post("/login", login.Login)
		l.Post("/register", login.Register)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		u.Post("/user", user.PostUser)
		u.Put("/user/{id}", user.UpdateUser)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		u.Get("/saldo/{username}", pay.FindByUsername)
		u.Put("/saldo", pay.UpdateSaldo)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		u.Get("/order/{username}", payHistory.FindByUsername)
		u.Get("/order/id/{order_id}", payHistory.FindByOrderId)
		u.Post("/order", payHistory.PayOrder)
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Println("Service running on " + host + ":" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Println("Error Starting Service")
	}
}
