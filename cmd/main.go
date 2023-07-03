package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "strings"
  "math"
  "runtime"
  "time"
  "strconv"
  "fmt"
  "log"
  "net/http"
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

// Example of a microbenchmark for string concatenation in Go (Echo)
func microbenchmark(c echo.Context) error {
  n, err := strconv.Atoi(c.QueryParam("n"))
  if err != nil {
    log.Fatal(err)
  }
  start := time.Now()
  var builder strings.Builder
  for i := 0; i < n; i++ {
    builder.WriteString(fmt.Sprintf("%v", i))
  }
  end := time.Now()
  return c.String(http.StatusOK, end.Sub(start).String())
}

type User struct {
  ID int `json:"id"`
  Name string `json:"name"`
}
  
func memoryTest(c echo.Context) error {
  n, err := strconv.Atoi(c.QueryParam("n"))
  if err != nil {
    log.Fatal(err)
  }
  startMem := runtime.MemStats{}
  runtime.ReadMemStats(&startMem)
  // start := time.Now()
  data := make([]User, n)
  for i := 0; i < n; i++ {
    data[i] = User{ID: i, Name: fmt.Sprintf("User %d", i)}
  }
  endMem := runtime.MemStats{}
  runtime.ReadMemStats(&endMem)
  return c.String(http.StatusOK, strconv.Itoa(int((endMem.Alloc-startMem.Alloc)/1024/1024)))
}

func cpuTest(c echo.Context) error {
  n, err := strconv.Atoi(c.QueryParam("n"))
  if err != nil {
    log.Fatal(err)
  }
  start := time.Now()
  var result float64
  for i := 0; i < n; i++ {
    result += math.Pow(float64(i), 2)
  }
  end := time.Now()
  return c.String(http.StatusOK, end.Sub(start).String())
}

func main() {
  e := echo.New()
  
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", home)
  e.GET("/microbenchmark", microbenchmark)
  e.GET("/memorytest", memoryTest)
  e.GET("/cputest", cpuTest)
  
  
  e.GET("/hello", hello)
  e.GET("/users", users)

  e.Start(":3000")
}