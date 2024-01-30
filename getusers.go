package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFlights(c *gin.Context) {
	rows, err := db.Query("SELECT flight_id, flight_name, flight_arriving_time, flight_departure_time, passenger_details, started_at, reached_at, is_active FROM flight")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var flights []Flight
	for rows.Next() {
		var flight Flight
		err := rows.Scan(&flight.FlightID, &flight.FlightName, &flight.FlightArrivingTime, &flight.FlightDepartureTime, &flight.PassengerDetails, &flight.StartedAt, &flight.ReachedAt, &flight.IsActive)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		flights = append(flights, flight)
	}
	c.JSON(http.StatusOK, flights)
}
