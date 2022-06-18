package main

import (
	"ransmart_pay/app/helper/helper"
	"ransmart_pay/app/server"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load("params/.env")
	helper.CheckEnv(err)

	// running server
	server.Execute()
}
