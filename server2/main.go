package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, GaoEy!")
	})

	done := make(chan bool)
	go func() {
		s := &http.Server{
			Addr: ":4444",
		}
		fmt.Println("Server exit:", e.StartServer(s))
		done <- true
	}()
	fmt.Println("Shutdown", e.Shutdown(context.Background()))
	<-done
}
