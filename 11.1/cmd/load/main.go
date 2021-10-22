package main

import (
	"fmt"
	"log"
	"store"
)

func main() {
	s := store.Open("test.bin")
	defer s.Close()
	var result int
	err := s.Load(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data:", result)
}
