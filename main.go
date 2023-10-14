package main

import (
	"log"
	"restaurant-api/api"
)

func main() {
	if err := api.StartApp(); err != nil {
		log.Fatal("there was an error inializing app", err.Error())
	}
}
