package version

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type versionResponse struct {
	Version string
}

func makeVersionEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return versionResponse{Version: svc.Version()}, nil
	}
}
