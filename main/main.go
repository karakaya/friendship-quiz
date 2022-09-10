package main

import (
	"log"

	"github.com/karakaya/friendship-quiz/main/app"
)

func main() {
	logger := log.Logger{}
	app := app.NewApp(logger, 80)

	app.Run()

}
