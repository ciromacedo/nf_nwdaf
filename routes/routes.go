package routes

import (
	"nf_nwdaf/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/nwdaf/api/v1/udr/collections", controllers.DisplayCollections)
}
