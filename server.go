package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/server-up", server)
	e.GET("/ping", ping)
	e.GET("/mamot", mamot)
	e.GET("/workshop", workshop)
	e.GET("/workshop-devops", workshop_devops)
	e.GET("/tokopedia-workshop", tokopedia_workshop)
	e.GET("/dimas", dimas)
	

	// Start serve
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}

func mamot(c echo.Context) error {
	return c.String(http.StatusOK, "Mamot was here")
}

func workshop(c echo.Context) error {
	return c.String(http.StatusOK, "Happy workshop")
}

func workshop_devops(c echo.Context) error {
	return c.String(http.StatusOK, "Happy workshop devops 2021")
}

func tokopedia_workshop(c echo.Context) error {
	return c.String(http.StatusOK, "Tokopedia devops workshop 2021, Enjoy the show")
}

func dimas(c echo.Context) error {
	return c.String(http.StatusOK, "Hallo Dimas")
}

func server(c echo.Context) error {
	file, err := os.Open("service_wake_date.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	datefile, err := ioutil.ReadAll(file)
	msg := "server up: " + string(datefile)
	return c.String(http.StatusOK, msg)
}
