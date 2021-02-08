package controllers

import (
	"encoding/json"
	"net/http"
	"nf_nwdaf/base"
	"nf_nwdaf/model"
)


func DisplayCollections(w http.ResponseWriter, r *http.Request){
	//recupera o nome de todas as coleções
	var CollectionsInfo []model.CollectionInfo
	CollectionsName := base.GetCollectionsName()
	for _, coll := range CollectionsName {
		qtd, _ := base.GetNumberOfRecordsInCollection(coll)
		CollectionsInfo = append(CollectionsInfo, model.CollectionInfo{NumberOfRecords: qtd, DocumentName: coll})
	}
	json.NewEncoder(w).Encode(CollectionsInfo)
}