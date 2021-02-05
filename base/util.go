package base

import (
	"encoding/json"
	"fmt"
	"nf_nwdaf/model"
	"os"
)

func GetConfiguration()model.Config{
	file, fail := os.Open("config/config.json")
	if fail != nil {
		panic(fail.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := model.Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err.Error())
	}
	return configuration
}

func GetServerPort()string{
	ConfigPort := GetConfiguration().Port
	return fmt.Sprintf("%s%d", ":", ConfigPort)
}
