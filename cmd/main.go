package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func home(c echo.Context) error {
  return c.String(200, "Hello World!")
}

func hello(c echo.Context) error {
  return c.String(200, "Hello!")
}

func users(c echo.Context) error {
  // Query database and return result
  return c.JSON(200, "[{name: john}, {name: jane}]")
}

func main() {
  e := echo.New()
  
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", home)
  e.GET("/hello", hello)
  e.GET("/users", users)

  e.Start(":3000")
}