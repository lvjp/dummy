package fortune

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/lvjp/dummy/pkg/gokitutils"
)

func MakeHandler(svc Service) http.Handler {
	createHandler := kithttp.NewServer(
		makeCreateEndpoint(svc),
		gokitutils.DecodeJsonRequest[createRequest],
		kithttp.EncodeJSONResponse,
	)
	readHandler := kithttp.NewServer(
		makeReadEndpoint(svc),
		gokitutils.DecodeJsonRequest[readRequest],
		kithttp.EncodeJSONResponse,
	)
	updateHandler := kithttp.NewServer(
		makeUpdateEndpoint(svc),
		gokitutils.DecodeJsonRequest[updateRequest],
		kithttp.EncodeJSONResponse,
	)
	deleteHandler := kithttp.NewServer(
		makeDeleteEndpoint(svc),
		gokitutils.DecodeJsonRequest[deleteRequest],
		kithttp.EncodeJSONResponse,
	)

	return &fortuneHandler{
		createHandler: createHandler,
		readHandler:   readHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

type fortuneHandler struct {
	createHandler http.Handler
	readHandler   http.Handler
	updateHandler http.Handler
	deleteHandler http.Handler
}

func (fh *fortuneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	switch r.Method {
	case http.MethodPut:
		handler = fh.createHandler
	case http.MethodGet:
		handler = fh.readHandler
	case http.MethodPost:
		handler = fh.updateHandler
	case http.MethodDelete:
		handler = fh.deleteHandler
	default:
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}
