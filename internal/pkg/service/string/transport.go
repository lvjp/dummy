package string

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/lvjp/dummy/pkg/gokitutils"
)

func MakeHandler(svc Service) http.Handler {
	uppercaseHandler := kithttp.NewServer(
		makeUppercaseEndpoint(svc),
		gokitutils.DecodeJsonRequest[uppercaseRequest],
		kithttp.EncodeJSONResponse,
	)
	countHandler := kithttp.NewServer(
		makeCountEndpoint(svc),
		gokitutils.DecodeJsonRequest[countRequest],
		kithttp.EncodeJSONResponse,
	)

	m := http.NewServeMux()
	m.Handle("/uppercase", uppercaseHandler)
	m.Handle("/count", countHandler)

	return m
}
