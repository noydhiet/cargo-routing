package transport

import (
	"cargo-tracking/datastruct"
	dt "cargo-tracking/datastruct"
	"cargo-tracking/service"
	"context"
	"encoding/json"
	"errors"

	//"go-kit/kit/endpoint"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

type GetStatus interface {
	GetStatusService(context.Context, dt.Delivery) []dt.Delivery
}

type getStatus struct{}

var ErrEmpty = errors.New("empty string")

func (getStatus) GetStatusService(_ context.Context, a dt.Delivery) []dt.Delivery {
	return call_ServiceGetStatus(a)
}

func call_ServiceGetStatus(a dt.Delivery) []dt.Delivery {

	message := service.GetStatus(a)

	return message
}

func makeGetStatusEndpoint(i GetStatus) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.GetStatusDeliveryRequest)
		paramDel := dt.Delivery{}
		paramDel.ROUTE_ID = req.ROUTE_ID
		paramDel.ITENARY_ID = req.ITENARY_ID
		i.GetStatusService(ctx, paramDel)
		return datastruct.GetStatusDeliveryResponse{
			ROUTE_ID:            1,
			ITENARY_ID:          1,
			ROUTING_STATUS:      "unload",
			TRANSPORT_STATUS:    "inport",
			LAST_KNOWN_LOCATION: "surabaya",
		}, nil
	}
}

func decodeGetStatusDeliveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetStatusDeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	a := getStatus{}

	GetStatusHandler := httptransport.NewServer(
		makeGetStatusEndpoint(a),
		decodeGetStatusDeliveryRequest,
		encodeResponse,
	)

	http.Handle("/GetStatusDelivery", GetStatusHandler)
}
