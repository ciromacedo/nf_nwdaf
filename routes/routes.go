package routes

import (
	"nf_nwdaf/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/udr/collections", controllers.DisplayCollections)
}
