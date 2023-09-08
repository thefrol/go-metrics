package main

import (
	"flag"
	"fmt"
	"github.com/Sofja96/go-metrics.git/internal/handlers"
	"github.com/Sofja96/go-metrics.git/internal/storage"
	"github.com/Sofja96/go-metrics.git/internal/utils"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

type Config struct {
	address string `env:"ADDRESS"`
}

var cfg Config

func main() {

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		cfg.address = envRunAddr
	} else {
		flag.StringVar(&cfg.address, "a", "localhost:8080", "address and port to run server")
		flag.Parse()
	}

	//utils.ParseFlags()
	s := storage.New()
	e := echo.New()
	e.GET("/", handlers.AllMetrics(s))
	e.GET("/value/:typeM/:nameM", handlers.ValueMetrics(s))
	e.POST("/update/:typeM/:nameM/:valueM", handlers.Webhook(s))
	fmt.Println("Running server on", utils.FlagRunAddr)
	err := e.Start(cfg.address)
	if err != nil {
		log.Fatal(err)
	}
}
