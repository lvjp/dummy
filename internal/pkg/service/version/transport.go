package version

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeVersionEndpoint(svc VersionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return versionResponse{Version: svc.Version()}, nil
	}
}

type versionResponse struct {
	Version string
}
