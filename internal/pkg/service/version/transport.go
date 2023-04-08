package version

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(svc Service) http.Handler {
	versionHandler := kithttp.NewServer(
		makeVersionEndpoint(svc),
		kithttp.NopRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	m := http.NewServeMux()
	m.Handle("/version", versionHandler)

	return m
}
