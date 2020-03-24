package gokit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	additionalMiddleware "github.com/batazor/shortlink/pkg/api/middleware"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeAddLinkEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(link.Link)
		if !ok {
			return nil, nil
		}

		responseCh := make(chan interface{})

		// TODO: send []byte format
		go notify.Publish(api_type.METHOD_ADD, req, responseCh, "RESPONSE_STORE_ADD")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			return nil, fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		case notify.Response:
			return r.Payload.(*link.Link), nil
		}

		return nil, nil
	}
}

func makeGetLinkEndpoint() endpoint.Endpoint {
	return func(_ context.Context, r interface{}) (interface{}, error) {
		vars := mux.Vars(r.(*http.Request))
		if vars["id"] == "" {
			return nil, errors.New(`{"error": "need set hash URL"}`)
		}

		// Parse request
		var request = &getRequest{
			Hash: vars["id"],
		}

		var (
			response     *link.Link
			responseLink ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_GET, request.Hash, responseCh, "RESPONSE_STORE_GET")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.(*link.Link) // nolint errcheck
			}
		}

		var errorLink *link.NotFoundError
		if errors.As(err, &errorLink) {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}
		if err != nil {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}

		responseLink = ResponseLink{
			response,
		}

		return responseLink, nil
	}
}

func makeGetListLinkEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var (
			response     []*link.Link
			responseLink []ResponseLink // for custom JSON parsing
			err          error
		)

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_LIST, nil, responseCh, "RESPONSE_STORE_LIST")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
			return nil, err
		case notify.Response:
			err = r.Error
			if err == nil {
				response = r.Payload.([]*link.Link) // nolint errcheck
			}
		}

		for l := range response {
			responseLink = append(responseLink, ResponseLink{response[l]})
		}

		return responseLink, nil
	}
}

func makeDeleteLinkEndpoint() endpoint.Endpoint {
	return func(_ context.Context, r interface{}) (interface{}, error) {
		var err error
		req, ok := r.(link.Link)
		if !ok {
			return nil, nil
		}

		responseCh := make(chan interface{})

		go notify.Publish(api_type.METHOD_DELETE, req.Hash, responseCh, "RESPONSE_STORE_DELETE")

		c := <-responseCh
		switch r := c.(type) {
		case nil:
			err = fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		case notify.Response:
			err = r.Error
		}

		if err != nil {
			return nil, errors.New(`{"error": "` + err.Error() + `"}`)
		}

		return `{}`, nil
	}
}

func (api API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer opentracing.Tracer) error { // nolint unparam
	api.ctx = ctx

	log.Info("Run go-kit API")

	linkAddHandler := httptransport.NewServer(
		makeAddLinkEndpoint(),
		decodeAddLinkRequest,
		encodeResponse,
	)

	linkGetByIdHandler := httptransport.NewServer(
		makeGetLinkEndpoint(),
		decodeGetLinkRequest,
		encodeResponse,
	)

	linkGetListHandler := httptransport.NewServer(
		makeGetListLinkEndpoint(),
		decodeGetLinkRequest,
		encodeResponse,
	)

	linkDeleteHandler := httptransport.NewServer(
		makeDeleteLinkEndpoint(),
		decodeAddLinkRequest,
		encodeResponse,
	)

	// set-up router and initialize http endpoints
	r := mux.NewRouter()

	r.Methods("GET").Path("/api/links").Handler(linkGetListHandler)
	r.Methods("GET").Path("/api/{id}").Handler(linkGetByIdHandler)
	r.Methods("POST").Path("/api").Handler(linkAddHandler)
	r.Methods("DELETE").Path("/api").Handler(linkDeleteHandler)

	// Additional middleware
	r.Use(additionalMiddleware.Logger(log))

	log.Info(fmt.Sprintf("Run on port %d", config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

	return err
}

func decodeAddLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request link.Link
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}