package datastruct

type GetStatusDeliveryRequest struct {
	ROUTE_ID   int `json:"route_id"`
	ITENARY_ID int `json:"itenary_id"`
	CARGO_ID   int `json:"cargo_id"`
	VOYAGE_ID  int `json:"voyage_id"`
}

type GetStatusDeliveryResponse struct {
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
	LASTKNOWN_LOC    string `json:"lastknown_loc"`
}

type UpdateStatusDeliveryRequest struct {
	ROUTE_ID         int    `json:"route_id"`
	ITENARY_ID       int    `json:"itenary_id"`
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
	LASTKNOWN_LOC    string `json:"lastknown_loc"`
}

type UpdateStatusDeliveryResponse struct {
	MESSAGE string `json:"message"`
}
