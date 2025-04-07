package main

import (
	"database/sql"
	"fmt"
	"internet-store/config"
	"internet-store/internal/repository"
	"internet-store/internal/server"
	"internet-store/internal/service"
	"log"

	_ "github.com/lib/pq"

	"github.com/caarlos0/env"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	usersRepository := repository.NewUserRepository(db)
	usersService := service.NewUserService(usersRepository)

	s := server.NewServer(usersService)

	e := s.SetupRouter()

	err = e.Run(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
