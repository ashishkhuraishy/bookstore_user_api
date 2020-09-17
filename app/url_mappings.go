package app

import (
	"github.com/ashishkhuraishy/bookstore_user_api/controllers/ping"
	"github.com/ashishkhuraishy/bookstore_user_api/controllers/user"
)

func urlMappings() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", user.GetUser)
	router.POST("/user/", user.Create)
}
