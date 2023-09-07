package main

import (
	"fmt"
	"github.com/Sofja96/go-metrics.git/internal/handlers"
	"github.com/Sofja96/go-metrics.git/internal/storage"
	"github.com/Sofja96/go-metrics.git/internal/utils"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	utils.ParseFlags()
	s := storage.New()
	e := echo.New()
	e.GET("/", handlers.AllMetrics(s))
	e.GET("/value/:typeM/:nameM", handlers.ValueMetrics(s))
	e.POST("/update/:typeM/:nameM/:valueM", handlers.Webhook(s))
	fmt.Println("Running server on", utils.FlagRunAddr)
	err := e.Start(utils.FlagRunAddr)
	if err != nil {
		log.Fatal(err)
	}
}
