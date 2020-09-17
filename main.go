package main

import (
	"github.com/ashishkhuraishy/bookstore_user_api/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.StartApp()
}
