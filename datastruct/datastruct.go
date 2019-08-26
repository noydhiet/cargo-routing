package datastruct

type GetStatusDeliveryRequest struct {
	ROUTE_ID    int `json:"route_id"`
	ITENARY_ID  int `json:"itenary_id"`
	DELIVERY_ID int `json:"delivery_id"`
}

type GetStatusDeliveryResponse struct {
	ROUTE_ID            int    `json:"route_id"`
	ITENARY_ID          int    `json:"itenary_id"`
	ROUTING_STATUS      string `json:"routing_status"`
	TRANSPORT_STATUS    string `json:"transport_status"`
	LAST_KNOWN_LOCATION string `json:"last_known_location"`
}

type UpdateStatusRequest struct {
	ROUTE_ID   int `json:"route_id"`
	ITENARY_ID int `json:"itenary_id"`
}

type UpdateStatusResponse struct {
	DELIVERY_ID      int    `json:"delivery_id"`
	ROUTE_ID         int    `json:"route_id"`
	ITENARY_ID       int    `json:"itenary_id"`
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
}

type Delivery struct { //using datastruct as an input also output
	DELIVERY_ID         int
	ROUTE_ID            int
	ITENARY_ID          int
	ROUTING_STATUS      string
	TRANSPORT_STATUS    string
	LAST_KNOWN_LOCATION string
}
