package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateFlight(c *gin.Context) {
	var u Flight
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	Flight_id := c.Param("flight_id")
	_, err := db.Exec("UPDATE flight SET flight_name = $1, flight_arriving_time = $2, flight_departure_time = $3, passenger_details = $4, started_at = $5, reached_at = $6, is_active = $7 WHERE flight_id = $8", u.FlightName, u.FlightArrivingTime, u.FlightDepartureTime, u.PassengerDetails, u.StartedAt, u.ReachedAt, u.IsActive, Flight_id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flight"})
		return
	}

	c.JSON(http.StatusOK, u)
}
