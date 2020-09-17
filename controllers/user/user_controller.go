package user

import (
	"fmt"
	"net/http"
	"strconv"

	service "github.com/ashishkhuraishy/bookstore_user_api/Services"
	"github.com/ashishkhuraishy/bookstore_user_api/domain"
	"github.com/ashishkhuraishy/bookstore_user_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// Create will create a new user
func Create(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequest("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := service.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser will get a user from the given id
func GetUser(c *gin.Context) {
	userParam := c.Param("user_id")
	userID, err := strconv.ParseInt(userParam, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest(fmt.Sprintf("%s is an invalid user id", userParam))
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := service.GetUser(userID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
