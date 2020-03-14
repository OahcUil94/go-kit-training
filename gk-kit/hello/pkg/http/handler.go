package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/OahcUil94/go-kit-training/gk-kit/hello/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeFooHandler creates the handler logic
func makeFooHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/foo", http1.NewServer(endpoints.FooEndpoint, decodeFooRequest, encodeFooResponse, options...))
}

// decodeFooRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeFooRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.FooRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeFooResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeFooResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
