package fortune

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type createRequest struct {
	Fortune string
}

type createResponse struct {
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}

func makeCreateEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		uuid, err := svc.Create(req.Fortune)
		if err != nil {
			return createResponse{"", err.Error()}, nil
		}
		return createResponse{uuid, ""}, nil
	}
}

type readRequest struct {
	Uuid string
}

type readResponse struct {
	Fortune string `json:"fortune,omitempty"`
	Error   string `json:"error,omitempty"`
}

func makeReadEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(readRequest)
		fortune, err := svc.Read(req.Uuid)
		if err != nil {
			return readResponse{"", err.Error()}, nil
		}
		return readResponse{fortune, ""}, nil
	}
}

type updateRequest struct {
	Uuid    string
	Fortune string
}

type updateResponse struct {
	Error string `json:"error,omitempty"`
}

func makeUpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := svc.Update(req.Uuid, req.Fortune)
		if err != nil {
			return updateResponse{err.Error()}, nil
		}
		return updateResponse{}, nil
	}
}

type deleteRequest struct {
	Uuid string
}

type deleteResponse struct {
	Error string `json:"error,omitempty"`
}

func makeDeleteEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		err := svc.Delete(req.Uuid)
		if err != nil {
			return deleteResponse{err.Error()}, nil
		}
		return deleteResponse{}, nil
	}
}
