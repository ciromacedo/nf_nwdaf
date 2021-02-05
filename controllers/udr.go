package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nf_nwdaf/model"
)

var Articles []model.Article
func DisplayCollections(w http.ResponseWriter, r *http.Request){
	 	Articles = []model.Article{
				model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
				model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}