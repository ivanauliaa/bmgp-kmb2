package main

import (
	"net/http"
	"strconv"

	formatter "github.com/ivanauliaa/response-formatter"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, formatter.BadRequestResponse(map[string]interface{}{"error": err.Error()}))
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, formatter.SuccessResponse(user))
		}
	}

	return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, formatter.BadRequestResponse(map[string]interface{}{"error": err.Error()}))
	}

	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, formatter.SuccessResponse(user))
		}
	}

	return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, formatter.BadRequestResponse(map[string]interface{}{"error": err.Error()}))
	}

	for index, user := range users {
		if user.Id == id {
			payload := User{}
			c.Bind(&payload)

			users[index].Name = payload.Name
			users[index].Email = payload.Email
			users[index].Password = payload.Password

			return c.JSON(http.StatusOK, formatter.SuccessResponse(users[index]))
		}
	}

	return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8100"))

	// API test
	// curl -X GET localhost:8100/users
	// curl -X POST -d '{"name": "Ivan", "email": "ivan@gmail.com", "password": "ivan"}' -H "Content-Type: application/json" localhost:8100/users
	// curl -X GET localhost:8100/users/1
	// curl -X PUT -d '{"name": "Ivan Aulia", "email": "ivan@gmail.com", "password": "ivan"}' -H "Content-Type: application/json" localhost:8100/users/1
	// curl -X DELETE localhost:8100/users/1
}
