package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFlightByID(c *gin.Context) {
	id := c.Param("id")

	var u Flight
	err := db.QueryRow("SELECT * FROM flight WHERE flight_id = $1", id).Scan(&u.FlightID, &u.FlightName, &u.FlightArrivingTime, &u.FlightDepartureTime, &u.PassengerDetails, &u.StartedAt, &u.ReachedAt, &u.IsActive)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
