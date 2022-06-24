package models

import "time"

type Endpoint struct {
	Url               string
	NumberOfHits      int
	LastThreeStatuses []Status
}

type Status struct {
	StatusCode int
	Time       time.Time
}

type JsonResponse struct {
	Url string `json:"url"`
}
