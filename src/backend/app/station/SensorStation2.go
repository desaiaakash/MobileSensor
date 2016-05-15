package station

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	)

	func Station2_Init() {

	mux := httprouter.New()
	mux.POST("/sensor/add/", AddSensor)			      								// Add  Sensor detail for sensors subscribed by given user
	mux.DELETE("/sensor/remove/:username/:pwd/:sensorid", RemoveSensor)		        // Remove Sensor detail for sensors subscribed by given user
	mux.GET("/sensor/status/:username/:pwd/:sensorid/:status", SubscribeSensor)    //Unsubscribe a given sensor-id by given user	
	mux.GET("/sensorlist/:username/:pwd/", getSensorInfo)                 		  // Retrieve Temperature/WaterLevel Sensor detail for sensors subscribed by given user
	mux.GET("/user/profile/:username/:pwd/",getUserProfile)						  	
	mux.PUT("/user/profile/:username/:pwd",editUserProfile)
	mux.GET("/user/billing/:username/:pwd",getBillingDetails)
	
	c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "DELETE","PUT"},
    AllowCredentials: true,
	})

	
	handler := c.Handler(mux)
	
	
	server:= http.Server {
		Addr:    "0.0.0.0:3502",
		Handler : handler,
		}
		
		server.ListenAndServe()
	}
		
		
