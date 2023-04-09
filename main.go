package main

import (
	"fmt"
	"io/ioutil"
	"log"

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
	repository.MongoConnectDatabase()

	app := routes.Initialize()

	printLogo()
	app.Listen(":8080")
}

func printLogo() {
	content, err := ioutil.ReadFile("./resources/logo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
