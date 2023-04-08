package debug

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

	buildTimestampHandler := kithttp.NewServer(
		makeBuildTimestampEndpoint(svc),
		kithttp.NopRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	environmentHandler := kithttp.NewServer(
		makeEnvironmentEndpoint(svc),
		kithttp.NopRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	m := http.NewServeMux()
	m.Handle("/version", versionHandler)
	m.Handle("/buildtimestamp", buildTimestampHandler)
	m.Handle("/environment", environmentHandler)

	return m
}
