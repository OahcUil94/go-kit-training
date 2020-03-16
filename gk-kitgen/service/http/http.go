package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/OahcUil94/go-kit-training/gk-kitgen/service/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/postprofile", httptransport.NewServer(endpoints.PostProfile, DecodePostProfileRequest, EncodePostProfileResponse))
	return m
}
func DecodePostProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.PostProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodePostProfileResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
