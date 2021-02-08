package main

import (
	"log"
	"os"
)

func main() {
	log.Println("init debug")
	log.Println(os.Getenv("DB_URI"))
	log.Println("end debug")
}
