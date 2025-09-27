package h

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Key   int `json:"key"`
	Value any `json:"value"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Query(r *http.Response, name string) string {
	return r.Request.URL.Query().Get(name)
}
