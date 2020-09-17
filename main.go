package main

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"github.com/edwinvautier/DB_VAUTIER_P01/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	// Connect to database
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Authorization",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.SetupRouter(r)

	log.Fatal(r.Run(":8000")) // listen and serve on 8000
}
