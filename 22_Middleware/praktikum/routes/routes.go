package routes

import (
	"praktikum/constants"
	c "praktikum/controllers"
	m "praktikum/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	m.LogMiddleware(e)

	e.POST("/auth/login", c.LoginController)

	e.POST("/users", c.CreateUserController)

	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)

	jwtAuth := e.Group("/restricted")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	jwtAuth.GET("/users", c.GetUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)

	jwtAuth.POST("/books", c.CreateBookController)
	jwtAuth.DELETE("/books/:id", c.DeleteBookController)
	jwtAuth.PUT("/books/:id", c.UpdateBookController)

	return e
}

// users example request
// curl -X GET localhost:8100/restricted/users
// curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/users

// curl -X POST -d '{"name": "James", "email": "james@gmail.com", "password": "james"}' -H "Content-Type: application/json" localhost:8100/users

// curl -X GET localhost:8100/restricted/users/3
// curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/users/3

// curl -X PUT -d '{"name": "James Rodriguez", "email": "james@gmail.com", "password": "james"}' -H "Content-Type: application/json" localhost:8100/restricted/users/3
// curl -X PUT -d '{"name": "James Rodriguez", "email": "james@gmail.com", "password": "james"}' -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/users/3

// curl -X DELETE localhost:8100/restricted/users/3
// curl -X DELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/users/3

// books example request
// curl -X GET localhost:8100/books

// curl -X POST -d '{"title": "Subtle of Not Giving a F", "year": "2018", "author": "Mark Manson"}' -H "Content-Type: application/json" localhost:8100/restricted/books
// curl -X POST -d '{"title": "Subtle of Not Giving a F", "year": "2018", "author": "Mark Manson"}' -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/books

// curl -X GET localhost:8100/books/2

// curl -X PUT -d '{"title": "James Rodriguez", "year": "2018", "author": "Mark Manson"}' -H "Content-Type: application/json" localhost:8100/restricted/books/2
// curl -X PUT -d '{"title": "James Rodriguez", "year": "2018", "author": "Mark Manson"}' -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/books/2

// curl -X DELETE localhost:8100/restricted/books/2
// curl -X DELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNjQ4NTg4NzI1LCJuYW1lIjoiSXZhbiBBdWxpYSIsInVzZXJJRCI6IjIifQ.XTLHBYvXXGDGxMMHUXrVBW6EGkt4LscQ-DotONUMVTM" localhost:8100/restricted/books/3
