package h

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func JSON(w http.ResponseWriter, status int, r Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
}

func Query(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}
