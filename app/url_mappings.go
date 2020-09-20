package app

import (
	"github.com/ashishkhuraishy/bookstore_user_api/controllers/ping"
	"github.com/ashishkhuraishy/bookstore_user_api/controllers/user"
)

func urlMappings() {
	// Ping to check if the server is alive
	router.GET("/ping", ping.Ping)

	// User Api
	router.POST("/user/", user.Create)

	// User Info Apis
	router.GET("/user/:user_id", user.GetUser)
	router.PUT("/user/:user_id", user.UpdateUser)
	router.PATCH("/user/:user_id", user.UpdateUser)
	router.DELETE("/user/:user_id", user.DeleteUser)

	// Search Apis
	// Required Query parameter `status`
	router.GET("internal/users/search/", user.Search)

}
