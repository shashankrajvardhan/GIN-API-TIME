package main

import (
	"github.com/gin-gonic/gin"
)

func router() {
	r := gin.Default()

	// Define routes for user API
	r.GET("/users", GetFlights)            
	r.GET("/getflight/:id", GetFlightByID) 
	r.POST("/createflight", CreateFlight)  
	r.PUT("/updateflight/:id", UpdateFlight)
	r.DELETE("/deleteflight/:id", DeleteFlight)
	r.GET("/flifgtname/:username", getbyname)
	r.GET("/flightpartial/:username", partial) 

	r.Run(":8080")
}
