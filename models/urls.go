package models

import "time"

//Endpoint struct holds url, previous statuses and two booleans, Blocked, this property says if
//endpoint is already served, PreviousStatuses, this property says the previous status returned.
type Endpoint struct {
	Url              string
	Blocked          bool
	IsReadyToServe   bool
	PreviousStatuses []Status
}

//Status struct holds the status code and the time the status returned.
type Status struct {
	StatusCode int
	Time       time.Time
}

type JsonResponse struct {
	Url string `json:"url"`
}
