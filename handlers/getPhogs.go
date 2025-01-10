package handlers

import "net/http"

func (cfg *ApiConfig) GetPhotographers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("photographers list"))
}
