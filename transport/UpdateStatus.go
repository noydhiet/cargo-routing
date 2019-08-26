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

type UpdateStatus interface {
	UpdateStatusService(context.Context, dt.Delivery) []dt.Delivery
}

type updateStatus struct{}

var ErrEmpty = errors.New("empty string")

func (updateStatus) UpdateStatusService(_ context.Context, a dt.Delivery) []dt.Delivery {
	return call_ServiceUpdateStatus(a)
}

func call_ServiceUpdateStatus(a dt.Delivery) []dt.Delivery {
	message := service.UpdateStatus(a)
	return message
}

func makeUpdateStatusEndpoint(i UpdateStatus) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.UpdateStatusRequest)
		paramDel := dt.Delivery{}
		paramDel.ROUTE_ID = req.ROUTE_ID
		paramDel.ITENARY_ID = req.ITENARY_ID
		i.UpdateStatusService(ctx, paramDel)

		return datastruct.UpdateStatusResponse{
			DELIVERY_ID:      4,
			ROUTE_ID:         1,
			ITENARY_ID:       1,
			ROUTING_STATUS:   "load",
			TRANSPORT_STATUS: "Docking",
		}, nil
	}
}

func decodeUpdateStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	x := updateStatus{}

	UpdateStatusHandler := httptransport.NewServer(
		makeUpdateStatusEndpoint(x),
		decodeUpdateStatusRequest,
		encodeResponse,
	)

	//make an endpoint
	http.Handle("/UpdateStatus", UpdateStatusHandler)
}
