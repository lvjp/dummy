package string

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewUppercaseHandler(svc StringService) http.Handler {
	return httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
}

func NewCountHandler(svc StringService) http.Handler {
	return httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
}
