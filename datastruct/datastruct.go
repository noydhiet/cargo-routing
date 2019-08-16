package datastruct

type GetStatusDeliveryRequest struct {
	ROUTE_ID   int `json:"route_id"`
	ITENARY_ID int `json:"itenary_id"`
}

type GetStatusDeliveryResponse struct {
	ROUTE_ID            int    `json:"route_id"`
	ITENARY_ID          int    `json:"itenary_id"`
	ROUTING_STATUS      string `json:"routing_status"`
	TRANSPORT_STATUS    string `json:"transport_status"`
	LAST_KNOWN_LOCATION string `json:"last_known_location"`
}

type UpdateStatusDeliveryRequest struct {
	ROUTE_ID   int `json:"route_id"`
	ITENARY_ID int `json:"itenary_id"`
}

type UpdateStatusDeliveryResponse struct {
	MESSAGE string `json:"message"`
}
