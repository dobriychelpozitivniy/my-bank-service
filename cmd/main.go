package main

import (
	"github.com/dobriychelpozitivniy/mybank"
	"github.com/dobriychelpozitivniy/mybank/pkg/handler"
	"github.com/dobriychelpozitivniy/mybank/pkg/repository"
	"github.com/dobriychelpozitivniy/mybank/pkg/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	db, err := repository.NewSqliteDB(repository.Config{
		DBName: viper.GetString("dbname"),
	})
	if err != nil {
		log.Fatalf("db error: %s", err.Error())
	}

	err = goose.SetDialect("sqlite3")
	if err != nil {
		log.Fatalf("dialect error: %s", err.Error())
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalf("migration error: %s", err.Error())
	}

	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(mybank.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("start server eror: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
