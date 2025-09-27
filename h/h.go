package h

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Key   int `json:"key"`
	Value any `json:"value"`
}

type Res func(string, any)

func JSON(w http.ResponseWriter, status int, Res Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Res)
}

func Query(r *http.Response, name string) string {
	return r.Request.URL.Query().Get(name)
}
