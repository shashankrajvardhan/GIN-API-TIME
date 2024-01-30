package main

import "time"

type Flight struct {
	FlightID            int       `json:"flight_id"`
	FlightName          string    `json:"flight_name"`
	FlightArrivingTime  time.Time `json:"flight_arriving_time"`
	FlightDepartureTime time.Time `json:"flight_departure_time"`
	PassengerDetails    string    `json:"passenger_details"`
	StartedAt           time.Time `json:"started_at"`
	ReachedAt           time.Time `json:"reached_at"`
	IsActive            bool      `json:"is_active"`
}
