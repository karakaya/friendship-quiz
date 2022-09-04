package main

import (
	"github.com/karakaya/friendship-quiz/pkg/db"
)

func main() {
	_, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

}
