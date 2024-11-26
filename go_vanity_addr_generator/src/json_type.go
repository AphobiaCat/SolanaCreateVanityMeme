package main

import (
	"encoding/json"
	//"reflect"
)

func Build_Json(obj interface{}) string {

	jsonData, err := json.Marshal(obj)
	if err != nil {
		DBG_ERR("Error marshalling JSON:", err)
	}

	return string(jsonData)
}
