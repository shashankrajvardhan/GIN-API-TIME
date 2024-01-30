package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFlight(c *gin.Context) {
	var u Flight
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the database query
	err := db.QueryRow("INSERT INTO flight(flight_name, passenger_details, is_active) VALUES ($1, $2, $3) RETURNING flight_id",
		u.FlightName, u.PassengerDetails, u.IsActive).Scan(&u.FlightID)

	if err != nil {
		log.Println("Failed to create flight:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}

	c.JSON(http.StatusCreated, u)
}
