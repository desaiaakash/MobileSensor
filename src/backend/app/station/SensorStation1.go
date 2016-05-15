package station

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"time"
	)

	type ResultSet struct{
	Result bool `json : "result"`	
	}

	type Users struct {
	Uid int `json:"userid"`
	Username   string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	}
	
	type Sensors struct {
	SensorId string `json:"id"`
	}

	type SensorInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	TypeCode int64`json:"typecode"`
	Value int    `json:"value"`
	Location   string `json:"location"`
	Color string `json:"color"`
	Status bool `json:"status"`
	Cost float64 `json:"cost"`
	Coordinate struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinate"`
	History []float64	`json:"history"`
	}

	type UserSensorDetails struct {
	SensorsInfo []SensorInfo
	}

	type AddSensorRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Location   string `json:"location"`
	Type       string `json:"type"`
	TypeCode   int32  `json:"typecode"`
	Coordinate struct {
		Lat float64 `json:"Lat"`
		Lng float64 `json:"Lng"`
	} `json:"Coordinate"`
	}

	type BillingDetails struct {
	 SensorId string `json:"id"`
	Location   string `json:"location"`
	SensorType string `json:"type"`
	TypeCode   int32  `json:"typecode"`
	Cost float64 `json:"cost"`
	TimeUsed float64 `json:"timeUsed"`
	}
	
	type BillingArray struct {
	Bills []BillingDetails
	}
	

	var db  *sql.DB
	var err error

	/*func init() {
    db, err = sql.Open("mysql", "root:root@/test")
    if err != nil {
        fmt.Println(err)
    }
    err = db.Ping()
          if err == nil {
              fmt.Println("DB Successfully connected!!!")
	      }
	}*/

	// webservice which retrieves sensor data for all sensors subscribed by given user from SensorDataGenerator
	//mux.GET("/sensorlist/:username/:pwd/", getSensorInfo)  
	func getSensorInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	err = row.Scan(&usr.Username)
	
	authenticate(err,w,r)


	rows, err := db.Query("SELECT sensorid FROM sensorusage where username =  ?", usr.Username)
	checkErrors(err)

	var sensorsInfo []SensorInfo
	var iSensorInfo SensorInfo

	for rows.Next() { // iterating rows from SQL query result

		iSensor := new(Sensors) //object of Sensor Struct

		err := rows.Scan(&iSensor.SensorId)
		checkErrors(err)

		urlFormat := "http://localhost:7890/sensor/data/current/" + iSensor.SensorId

		fmt.Println("Starting....")
		fmt.Println(urlFormat)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", urlFormat, nil)
		resp, _ := client.Do(req)

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			body, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(body, &iSensorInfo)
		}

		sensorsInfo = append(sensorsInfo, iSensorInfo)
	}
	

	result := UserSensorDetails{sensorsInfo}
	resJson, _ := json.Marshal(result)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resJson)

	}


	
	// webservice which adds a new sensor and receives sensor data from SensorDataGenerator
	func AddSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params)   { 
		
	var timeUsed float64
	a :=  AddSensorRequest{}
	json.NewDecoder(r.Body).Decode(&a)
	
	username:= a.Username
	password:= a.Password
		 
	a.TypeCode = getTypeCode(a.Type)
		
	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	err = row.Scan(&usr.Username)
	authenticate(err,w,r)
	
	
	urlStr := "http://localhost:7890/sensor/add"
	fmt.Println("Testing.."+urlStr)
	buf, err := json.Marshal(a)
	 
	if err != nil {
  		fmt.Println("Marshalling failed:", err)
  	} else if err == nil {
  		fmt.Println("Marshalling successfull")
  	}
 

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(buf))
	
	req.Header.Set("X-Custom-Header", "SensorInformation")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	loc, _ := time.LoadLocation("UTC")
	startTime:= time.Now().In(loc)

	resp, err := client.Do(req)
	checkErrors(err)

	endTime:= time.Now().In(loc)
	timeDiff := endTime.Sub(startTime)
	timeUsed = timeDiff.Seconds()
	fmt.Println("Testing time Difference : ",timeUsed)
	current_date:= time.Now().Local().Format("2006-01-02")
	defer resp.Body.Close()
		 
	body,err := ioutil.ReadAll(resp.Body)
	checkErrors(err)

	
	var m SensorInfo
	json.Unmarshal(body,&m)
	
	_, err= db.Exec("insert into sensorusage (username,sensorid,start_time,end_time,time_used) values (?, ?, ?,?,?)",username,m.Id,startTime,endTime,timeUsed)
	checkErrors(err)
	
	_, err= db.Exec("insert into sensorinfo (sensorid,sensortype,sensorlocation,created_date,cost,status) values (?,?,?,?,?,?)",m.Id,m.Type,m.Location,current_date,m.Cost,m.Status)
	checkErrors(err)
	
	resJson, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	
   	}

	
	// webservice which remove a particular sensor
	func RemoveSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	fmt.Println("Calling remove sensor")
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	sensorid:= p.ByName("sensorid")
	fmt.Println(username,password)

	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	err = row.Scan(&usr.Username)
	authenticate(err,w,r)
	t:= time.Now().Local().Format("2006-01-02")
	_, err = db.Exec("Update sensorinfo SET end_date = ?,status = 0 WHERE sensorid = ?" ,t,sensorid)
	checkErrors(err)
	GetResultJson(err,w,r)
	
	
	}
	
	// webservice which changes the subscription status of a particular sensor
	func SubscribeSensor(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
		
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	sensorid:= p.ByName("sensorid")
	status:= p.ByName("status")

	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	err = row.Scan(&usr.Username)
	authenticate(err,w,r)
	
	// update sensor status to subscribe/unsubscribe
    _,err = db.Exec("update sensorinfo set status=? where sensorid=?",status,sensorid)
    checkErrors(err)
    GetResultJson(err,w,r)
    
	}
	
	


	// webservice which retrieves the profile of a particular sensor
	//mux.GET("/user/profile/:username/:pwd/",getUserProfile)
	func getUserProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	fmt.Println(username,password)

	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)

	err = row.Scan(&usr.Username)
	authenticate(err,w,r)

	row = db.QueryRow("SELECT userid,password,FirstName,LastName,Address,City,State,Country FROM users WHERE username = ?", username)
	err = row.Scan(&usr.Uid,&usr.Password,&usr.FirstName,&usr.LastName,&usr.Address,&usr.City,&usr.State,&usr.Country)
	checkErrors(err)
	
	resJson, _ := json.Marshal(usr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	
	}

	// webservice which updates the profile of a particular sensor
	// mux.PUT("/user/profile/:username/:pwd",editUserProfile)
	func editUserProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	row:= db.QueryRow("SELECT userid,username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	var a Users
	err = row.Scan(&usr.Uid,&usr.Username)
	authenticate(err,w,r)

	json.NewDecoder(r.Body).Decode(&a)
	a.Uid = usr.Uid
	a.Username = usr.Username
	fmt.Println(a.Password,a.FirstName,a.LastName,a.Address,a.City,a.State,a.Country,a.Username)
	_, _ = db.Exec("update users SET password = ?,FirstName = ?,LastName = ?,Address = ?, City = ?,State = ?,Country = ? WHERE username = ?" , a.Password,a.FirstName,a.LastName,a.Address,a.City,a.State,a.Country,a.Username)

	checkErrors(err)
	//update users SET password = pass1234,FirstName = Vignesh,LastName = Ramkumar,Address = '101 E San Fernando St', City = 'San Jose',State = California,Country = US WHERE username = 'abc@gmail.com';
	
	resJson, _ := json.Marshal(a)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)

	}

	// webservice which gets Billing details of a particular person
	//mux.GET("/user/billing/:username/:pwd",getBillingDetails)
	func getBillingDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) { 
	
	username:= p.ByName("username")
	password:= p.ByName("pwd")
	
	row:= db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username,password)
	
	usr:= new(Users)
	//var a Users
	err = row.Scan(&usr.Username)
	authenticate(err,w,r)

	rows, err := db.Query("select A.sensorid,A.sensorlocation,A.sensortype,A.cost,B.time_used from sensorinfo A inner join sensorusage B on A.sensorid = B.sensorid where username =  ?", usr.Username)
	checkErrors(err)

	var billingDetails []BillingDetails
	var iBillingDetails BillingDetails
	
	for rows.Next() { // iterating rows from SQL query result
		
		err := rows.Scan(&iBillingDetails.SensorId,&iBillingDetails.Location,&iBillingDetails.SensorType,&iBillingDetails.Cost,&iBillingDetails.TimeUsed)
		checkErrors(err)
		
		iBillingDetails.TypeCode = getTypeCode(iBillingDetails.SensorType)
		billingDetails = append(billingDetails, iBillingDetails)
	}

	resJson, _ := json.Marshal(billingDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", resJson)
	}

	//method to check for SQL errors
	func checkErrors(err error) {
	if err != nil {
        panic(err)
	}
	}

	//method to authenticate a user in order to use web-service
	func authenticate(err error,w http.ResponseWriter, r *http.Request) {
	if err == nil {
	fmt.Println("User login successfull")
	} else if err == sql.ErrNoRows {
	fmt.Println("User Not Authenticated")
	http.NotFound(w, r)
	log.Fatal(err)
	}	
	}

	//method to retrieve Sensor TypeCode given an input SensorType
	func getTypeCode(SensorType string) int32 {
	
	switch SensorType {
	case "Temperature": 
	return 1
	case "Pressure": 
	return 2
	case "Humidity": 
	return 3
	case "Air Quality Index": 
	return 4
	default: return 0
		}
	}

	//method to print status of response from webservice based on success/failure
	func GetResultJson(err error,w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	if err == nil {
	result.Result = true;
	res,_ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", res)
	} else if err != nil {
		result.Result = false;
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	}	
	}
	
	func Station1_Init() {

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
		Addr:    "0.0.0.0:3501",
		Handler : handler,
		}
		
		server.ListenAndServe()
	}
		
		
