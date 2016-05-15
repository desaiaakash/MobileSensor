/**

Types of sensors we are considering -

1. Temperature (#0080ff) (1)
2. Pressure (#ff4000) (2)
3. Humidity (#4000ff) (3)
4. Air Quality Index (#39e600) (4)

Each sensor will give value for any one type. Each time webservice is called,
updated value is fetched from the server.
*/

package sensor

import "github.com/gorilla/mux"
import "net/http"
import "encoding/json"
import "strconv"
import "fmt"

var serverport int

func StartServer(port int) {

	fmt.Println("Starting Sensor Data Server : " + strconv.Itoa(port))
	serverport = port

	rtr := mux.NewRouter()
	rtr.HandleFunc("/sensor/data/current/{sensorid}", getSensorDataService).Methods("GET")
	rtr.HandleFunc("/sensor/add", addSensorService).Methods("POST")
	rtr.HandleFunc("/sensor/remove/{sensorid}", removeSensorService).Methods("DELETE")
	rtr.HandleFunc("/sensor/status/{sensorid}/{status}", changeSensorStatusService).Methods("PUT")
	rtr.HandleFunc("/sensor/list", getSensorListService).Methods("POST")
	http.Handle("/", rtr)
	http.ListenAndServe(":"+strconv.Itoa(serverport), nil)

}

// webservice which adds a new sensor and return sensor data
func addSensorService(rw http.ResponseWriter, req *http.Request) {

	//creating an empty input struct
	asreq := AddSensorReq{}

	// Populate the user data
	json.NewDecoder(req.Body).Decode(&asreq)

	res := addSensor(asreq)

	resJson, _ := json.Marshal(res)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", resJson)
}

// webservice which gets sensor data for a particular sensor id
func getSensorDataService(rw http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	sensorId := p["sensorid"]

	res := getSensorValue(sensorId)

	resJson, _ := json.Marshal(res)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resJson)
}

// webservice which removes sensor
func removeSensorService(rw http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	sensorId := p["sensorid"]

	removeSensor(sensorId)

	rw.WriteHeader(204)
}

// webservice which sets sensor status for a particular sensor id
func changeSensorStatusService(rw http.ResponseWriter, req *http.Request) {

	p := mux.Vars(req)

	sensorId := p["sensorid"]
	sensorStatus, _ := strconv.ParseBool(p["status"])

	res := changeSensorStatus(sensorId, sensorStatus)

	resJson, _ := json.Marshal(res)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resJson)
}

// webservice which returns sensor data for different sensors
//takes an array fo sensor ids and returns a list of sensors
func getSensorListService(rw http.ResponseWriter, req *http.Request) {

	slreq := SensorListReq{}

	json.NewDecoder(req.Body).Decode(&slreq)

	res := getSensorList(slreq)

	resJson, _ := json.Marshal(res)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resJson)
}
