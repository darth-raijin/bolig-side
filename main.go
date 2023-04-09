package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/darth-raijin/bolig-side/api/routes"
	_ "github.com/darth-raijin/bolig-side/docs"
	"github.com/joho/godotenv"
)

var (
	Version string = "1.0.0"
)

// @title bolig-side
// @description REST API server for bolig-side aka 'the Feedback' app
func main() {
	loadEnvironmentConfig()

	app := routes.Initialize()

	printLogo()
	app.Listen(":8080")
}

func loadEnvironmentConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

}

func printLogo() {
	content, err := ioutil.ReadFile("./resources/logo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
