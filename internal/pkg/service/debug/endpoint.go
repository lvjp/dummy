package debug

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

type buildTimestampResponse struct {
	BuildTimestamp string
}

func makeBuildTimestampEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return buildTimestampResponse{BuildTimestamp: svc.BuildTimestamp()}, nil
	}
}

type environmentResponse struct {
	Environment []string
}

func makeEnvironmentEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return environmentResponse{Environment: svc.Environment()}, nil
	}
}
