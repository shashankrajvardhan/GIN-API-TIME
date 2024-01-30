package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getbyname(c *gin.Context) {
	flight_name := c.Param("flight_name")

	var u Flight
	err := db.QueryRow("SELECT * FROM flight WHERE flight_id = $1", flight_name).Scan(&u.FlightID, &u.FlightName, &u.FlightArrivingTime, &u.FlightDepartureTime, &u.PassengerDetails, &u.StartedAt, &u.ReachedAt, &u.IsActive)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}
