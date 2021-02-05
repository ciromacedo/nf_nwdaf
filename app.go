package main

import (
	"fmt"
	"log"
	"net/http"
	"nf_nwdaf/base"
	"nf_nwdaf/routes"
)

func main() {
	routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(base.GetServerPort(), nil))
	fmt.Println("NWDAF Server Start!")
}
