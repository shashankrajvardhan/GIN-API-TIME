package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFlight(c *gin.Context) {
	id := c.Param("id")

	var u Flight
	err := db.QueryRow("SELECT * FROM flight WHERE flight_id = $1", id).Scan(&u.FlightID, &u.FlightName, &u.FlightArrivingTime, &u.FlightDepartureTime, &u.PassengerDetails, &u.StartedAt, &u.ReachedAt, &u.IsActive)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else {
		_, err := db.Exec("DELETE FROM flight WHERE flight_id = $1", id)
		if err != nil {
			// Handle the error appropriately
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
