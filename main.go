package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println("hell buddy")

	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	// server := NewAPIServer(":3000")

	// server.Run()
}
