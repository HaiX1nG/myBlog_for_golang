package main

import (
	"awesomeProject/src/router"
	"log"
)

func main() {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		log.Println("Service start Error: ", err)
	}
}
