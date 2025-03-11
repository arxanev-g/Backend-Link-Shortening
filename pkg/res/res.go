package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, satusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(satusCode)
	json.NewEncoder(w).Encode(data)
}
