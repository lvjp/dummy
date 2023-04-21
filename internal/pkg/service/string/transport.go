package string

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(svc Service) http.Handler {
	uppercaseHandler := kithttp.NewServer(
		makeUppercaseEndpoint(svc),
		decodeJsonRequest[uppercaseRequest],
		kithttp.EncodeJSONResponse,
	)
	countHandler := kithttp.NewServer(
		makeCountEndpoint(svc),
		decodeJsonRequest[countRequest],
		kithttp.EncodeJSONResponse,
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
