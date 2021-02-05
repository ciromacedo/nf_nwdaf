package model

type Config struct{
	Port int
	MongoURI string
	DBName string
}

type CollectionInfo struct {
	DocumentName string  `json:"Name"`
	NumberOfRecords int64
}

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}