package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/darth-raijin/bolig-side/api/routes"
	_ "github.com/darth-raijin/bolig-side/docs"
	"github.com/darth-raijin/bolig-side/pkg/repository"
	"github.com/darth-raijin/bolig-side/pkg/utility"
)

var (
	Version string = "1.0.0"
)

// @title bolig-side
// @description REST API server for bolig-side aka 'the Feedback' app
func main() {
	utility.LoadConfig()
	verifyDatabaseConnection()

	app := routes.Initialize()

	printLogo()
	app.Listen(":8080")
}

func verifyDatabaseConnection() {
	for {
		client, _, connectErr := repository.MongoConnectDatabase("Users")

		if connectErr == nil {
			utility.Log(utility.INFO, "Succesfully connected to database")
			repository.DisconnectMongo(context.Background(), client)
			break
		}

		if connectErr != nil {
			utility.Log(utility.ERROR, fmt.Sprintf("Failed to connect to database: %s", connectErr.Error()))
			if client != nil {
				repository.DisconnectMongo(context.Background(), client)
			}
			time.Sleep(10 * time.Second)

		}
	}
}

func printLogo() {
	content, err := ioutil.ReadFile("./resources/logo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
