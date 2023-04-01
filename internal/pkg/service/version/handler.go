package version

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewVersionHandler(svc VersionService) http.Handler {
	return httptransport.NewServer(
		makeVersionEndpoint(svc),
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
	)
}
