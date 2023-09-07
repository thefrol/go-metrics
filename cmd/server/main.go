package main

import (
	"github.com/Sofja96/go-metrics.git/internal/handlers"
	"github.com/Sofja96/go-metrics.git/internal/storage"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	s := storage.New()
	e := echo.New()
	e.GET("/", handlers.AllMetrics(s))
	e.GET("/value/:typeM/:nameM", handlers.ValueMetrics(s))
	e.POST("/update/:typeM/:nameM/:valueM", handlers.Webhook(s))
	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
