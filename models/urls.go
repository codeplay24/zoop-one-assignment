package models

import "time"

type Endpoint struct {
	Url              string
	Blocked          bool
	IsReadyToServe   bool
	PreviousStatuses []Status
}

type Status struct {
	StatusCode int
	Time       time.Time
}

type JsonResponse struct {
	Url string `json:"url"`
}
