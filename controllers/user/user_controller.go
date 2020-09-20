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

	c.JSON(http.StatusOK, result.Marshaller(checkPublic(c)))
}

// UpdateUser : updates a user with a specific userId
func UpdateUser(c *gin.Context) {
	userParam := c.Param("user_id")
	userID, err := strconv.ParseInt(userParam, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest(fmt.Sprintf("%s is an invalid user id", userParam))
		c.JSON(restErr.Status, restErr)
		return
	}

	var user domain.User
	err = c.BindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	result, restErr := service.UpdateUser(user, c.Request.Method == http.MethodPatch)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshaller(checkPublic(c)))
}

// DeleteUser : Endpoint to Delete a user with an id
func DeleteUser(c *gin.Context) {
	userParam := c.Param("user_id")
	userID, err := strconv.ParseInt(userParam, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest(fmt.Sprintf("%s is an invalid userID", userParam))
		c.JSON(restErr.Status, restErr)
		return
	}

	restErr := service.DeleteUser(userID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search will search based on the given query parameter
func Search(c *gin.Context) {
	statusQuery := c.Query("status")

	if statusQuery == "" {
		restErr := errors.NewBadRequest("query parameter `status` should not be empty")
		c.JSON(restErr.Status, restErr)
		return
	}

	users, err := service.Search(statusQuery)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshaller(checkPublic(c)))

}

// Helper function to check if the request contains header for public
func checkPublic(c *gin.Context) bool {
	return c.Request.Header.Get("X-Public") == "true"
}
