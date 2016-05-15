package main

import (
	
	"github.com/julienschmidt/httprouter"
	"net/http"
	"backend/app/station"
	"backend/app/loadbalancer"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
	)

	func init() {
		go station.Authorization()
		go station.Station1_Init()
		go station.Station2_Init()
	}

	
	

	func addSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params)   { 
	
	url:= loadbalancer.GetDns()
	url =  "http://" + url + "/sensor/add/"
	
	client := &http.Client{}
	
	a :=  station.AddSensorRequest{}
	json.NewDecoder(r.Body).Decode(&a)
	buf, _ := json.Marshal(a)
	

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	response, err := client.Do(request)
	

	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	fmt.Println("Load balancing for addSensor request. Finding the idle server")
	body,_:= ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	var m station.SensorInfo
	json.Unmarshal(body,&m)
	resJson, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	}

	func removeSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	url:= loadbalancer.GetDns()

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	sensorid:= p.ByName("sensorid")

	url =  "http://" + url + "/sensor/remove/" + username + "/" + password + "/" + sensorid
	fmt.Println(url)
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", url, nil)
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	_,err = ioutil.ReadAll(response.Body)
	station.GetResultJson(err,w,r)
	}

	func SubscribeSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	url:= loadbalancer.GetDns()

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	sensorid:= p.ByName("sensorid")
	status:= p.ByName("status")

	url =  "http://" + url + "/sensor/status/" + username + "/" + password + "/" + sensorid + "/" + status
	
	fmt.Println(url)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	_,err = ioutil.ReadAll(response.Body)
	station.GetResultJson(err,w,r)
	}

	
	func getSensorInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	url:= loadbalancer.GetDns()
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	url =  "http://" + url + "/sensorlist/" + username + "/" + password + "/" 
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	body,err := ioutil.ReadAll(response.Body)
	var m station.UserSensorDetails
	json.Unmarshal(body,&m)
	resJson, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	}

	func getUserProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	url:= loadbalancer.GetDns()

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	url =  "http://" + url + "/user/profile/" + username + "/" + password + "/" 
	fmt.Println(url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	body,err := ioutil.ReadAll(response.Body)
	
	fmt.Println(string(body))
	
	var m station.Users

	json.Unmarshal(body,&m)
	resJson, _ := json.Marshal(m)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	}

	func editUserProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	url:= loadbalancer.GetDns()

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	var a station.Users
	json.NewDecoder(r.Body).Decode(&a)
	buf, _ := json.Marshal(a)
	

	url =  "http://" + url + "/user/profile/" + username + "/" + password + "/" 
	fmt.Println(url)

	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(buf))
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	_,err = ioutil.ReadAll(response.Body)
	station.GetResultJson(err,w,r)
	
	}

	func getBillingDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	url:= loadbalancer.GetDns()

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	/*var a station.BillingDetails
	json.NewDecoder(r.Body).Decode(&a)
	*/
	
	url =  "http://" + url + "/user/billing/" + username + "/" + password + "/" 
	fmt.Println(url)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	
	if err != nil {
    fmt.Println(err)
	}

	defer response.Body.Close()
	
	body,err := ioutil.ReadAll(response.Body)
	var m station.BillingDetails
	json.Unmarshal(body,&m)
	resJson, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)
	
	}
	func main() {

	mux := httprouter.New()
	mux.POST("/sensor/add/", addSensor)			      								// Add  Sensor detail for sensors subscribed by given user
	mux.DELETE("/sensor/remove/:username/:pwd/:sensorid", removeSensor)		    // Remove Sensor detail for sensors subscribed by given user
	mux.GET("/sensor/status/:username/:pwd/:sensorid/:status", SubscribeSensor)    //Unsubscribe a given sensor-id by given user	
	mux.GET("/sensorlist/:username/:pwd/", getSensorInfo)                 		  // Retrieve Temperature/WaterLevel Sensor detail for sensors subscribed by given user
	mux.GET("/user/profile/:username/:pwd/",getUserProfile)						  	
	mux.PUT("/user/profile/:username/:pwd",editUserProfile)
	mux.GET("/user/billing/:username/:pwd",getBillingDetails)
	
	server := http.Server{
		Addr:    "0.0.0.0:3500",
		Handler: mux,
	}
	server.ListenAndServe()

}

