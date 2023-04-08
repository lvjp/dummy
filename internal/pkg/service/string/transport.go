package string

import (
	"context"
	"encoding/json"
	"net/http"

	pouetkithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(svc Service) http.Handler {
	uppercaseHandler := pouetkithttp.NewServer(
		makeUppercaseEndpoint(svc),
		decodeJsonRequest[uppercaseRequest],
		pouetkithttp.EncodeJSONResponse,
	)
	countHandler := pouetkithttp.NewServer(
		makeCountEndpoint(svc),
		decodeJsonRequest[countRequest],
		pouetkithttp.EncodeJSONResponse,
	)

	m := http.NewServeMux()
	m.Handle("/uppercase", uppercaseHandler)
	m.Handle("/count", countHandler)

	return m
}

func decodeJsonRequest[T interface{}](_ context.Context, r *http.Request) (interface{}, error) {
	var request T
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
