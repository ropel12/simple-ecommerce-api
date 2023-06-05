package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/ropel12/simple-ecommerce-api/config"
	"github.com/ropel12/simple-ecommerce-api/controller"
	"github.com/ropel12/simple-ecommerce-api/db"
	"github.com/ropel12/simple-ecommerce-api/helper"
	"github.com/ropel12/simple-ecommerce-api/middleware"
	"github.com/ropel12/simple-ecommerce-api/repository"
	"github.com/ropel12/simple-ecommerce-api/router"
	"github.com/ropel12/simple-ecommerce-api/service"
)

func main() {
	err := db.Newmigrate()
	helper.PanicIfError(err)
	db := db.NewDB()
	defer db.Close()
	validate := validator.New()
	productrepo := repository.NewProductRepo()
	Transactionrepo := repository.NewTransactionRepository()
	UserRepo := repository.NewUserRepository()
	PaymentRepo := repository.NewPaymentRepo()
	router := router.NewRouter(controller.NewInitControler(service.RunService(db, validate, UserRepo, productrepo, Transactionrepo, PaymentRepo)))
	server := http.Server{
		Addr:    "localhost:" + config.PORT,
		Handler: middleware.AuthtenticationMiddleware(router),
	}
	server.ListenAndServe()

}
