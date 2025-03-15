package model

type ProductListResponse struct {
	Data []Product `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ActionResponse struct {
	Message string `json:"message"`
}

type StatPercentResponse struct {
	Data []map[string]int `json:"data"`
}

type DistanceResponse struct {
	Data Distance `json:"data"`
}

type Distance struct {
	DistanceKm float64 `json:"distance_km"`
}
