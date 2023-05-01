package gokitutils

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeJsonRequest[T any](_ context.Context, r *http.Request) (interface{}, error) {
	var request T
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
