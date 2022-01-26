package user

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(endpoint *Endpoints) http.Handler {
	router := mux.NewRouter()

	router.Methods("POST").Path("/api").Handler(transport.NewServer(
		endpoint.CreateUser,
		createDecodeReq,
		encodeCreateUserResp,
	),
	)
	return router
}

func createDecodeReq(_ context.Context, r *http.Request) (interface{}, error) {

	var request CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCreateUserResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
