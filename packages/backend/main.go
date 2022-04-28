package main

import (
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/klyngen/starbook/api"
	"github.com/klyngen/starbook/person"
	"github.com/klyngen/starbook/star"
)

type configuration struct {
	Database string
	Username string
	Password string
	Host     string
}

func (c *configuration) isValid() bool {
	return len(c.Database) > 2 && len(c.Database) > 1 && len(c.Password) > 2 && len(c.Host) > 2
}

func main() {

	config := getConfiguration()

	if !config.isValid() {
		log.Panic("No valid configuration was provided. Please set the Database, Username, Password and Host")
	}

	log.Printf("Fetched configuration:\n\tUsername: %v\n\tDatabase: %v\n\tHost: %v", config.Username, config.Database, config.Host)

	starRepo, err := star.NewGormStarRepository(config.Host, config.Username, config.Password, config.Database)

	if err != nil {
		log.Fatalf("Unable to create a database connection: \n %v", err)
		log.Panic("Exiting")
	}
	log.Println("Created Star Repository")

	personRepo, err := person.NewGormPersonRepository(config.Host, config.Username, config.Password, config.Database)

	if err != nil {
		log.Fatalf("Unable to create a database connection: \n %v", err)
		log.Panic("Exiting")
	}

	log.Println("Created Person Repository")

	title := figure.NewFigure("Starbook V1.0", "", true)
	title.Print()

	log.Println("Starting server on port 8080")

	server := api.NewPresentationLayer(personRepo, starRepo)

	server.StartServer("8080")
}

func getConfiguration() configuration {
	return configuration{
		Database: os.Getenv("DATABASE_NAME"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
	}
}
