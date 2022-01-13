package user

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rafaeldiazmiles/ProjectEssay/pkg/user"
)

func NewHTTPServer(endpoint user.Endpoints) http.Handler {
	router := mux.NewRouter()

	router.Methods("POST").Path("/api").Handler(transport.NewServer(
		endpoint.CreateUser,
		createDecodeReq,
		encodeCreateUserResp,
	),
	)
}

func createDecodeReq(_ context.Context, r *http.Request) (interface{}, error) {

	var request user.CreateUserRequest
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
