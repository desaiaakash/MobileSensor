package main

//import "backend/sensor"
import "fmt"
import "github.com/gorilla/mux"
import "net/http"

//import "encoding/json"
import "strconv"
import "io/ioutil"
import "strings"

//import "encoding/json"

var port int = 6573
var serverport1 int = 6577
var serverport2 int = 7329
var count int = 0

func main() {
	fmt.Println("Starting Load Balancer : " + strconv.Itoa(port))
	startLBServer(port)
}

func startLBServer(port int) {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/sensor/data/current/{sensorid}", getSensorDataService).Methods("GET")
	rtr.HandleFunc("/sensor/add", addSensorService).Methods("POST")
	rtr.HandleFunc("/sensor/remove/{sensorid}", removeSensorService).Methods("DELETE")
	// rtr.HandleFunc("/sensor/status/{sensorid}/{status}", changeSensorStatusService).Methods("PUT")
	http.Handle("/", rtr)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// webservice which adds a new sensor and return sensor data
func addSensorService(rw http.ResponseWriter, req *http.Request) {

	port := strconv.Itoa(getServerPort())
	url := "http://localhost:" + port + "/sensor/add"

	//get json
	htmlReqByteArray, _ := ioutil.ReadAll(req.Body)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, strings.NewReader(string(htmlReqByteArray)))
	response, _ := client.Do(request)

	//get json
	htmlResByteArray, _ := ioutil.ReadAll(response.Body)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", htmlResByteArray)
}

// webservice which gets sensor data for a particular sensor id
func getSensorDataService(rw http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	sensorid := p["sensorid"]

	port := strconv.Itoa(getServerPort())
	url := "http://localhost:" + port + "/sensor/data/current/" + sensorid

	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(request)

	//convert body to json
	htmlDataByteArray, _ := ioutil.ReadAll(response.Body)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", htmlDataByteArray)
}

//webservice which removes sensor
func removeSensorService(rw http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	sensorId := p["sensorid"]

	port := strconv.Itoa(getServerPort())
	url := "http://localhost:" + port + "/sensor/remove/" + sensorId

	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", url, nil)
	client.Do(request)

	rw.WriteHeader(204)
}

// //webservice which sets sensor status for a particular sensor id
// func changeSensorStatusService(rw http.ResponseWriter, req *http.Request) {

// 	p := mux.Vars(req)

// 	sensorId := p["sensorid"]
// 	sensorStatus, _ := strconv.ParseBool(p["status"])

// 	res := changeSensorStatus(sensorId, sensorStatus)

// 	resJson, _ := json.Marshal(res)

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(200)
// 	fmt.Fprintf(rw, "%s", resJson)
// }

func getServerPort() int {
	count++

	if count%2 == 0 && count%5 == 0 {
		fmt.Println("Request handled by Sensor Server 1")
		return serverport1
	} else if count%2 != 0 {
		fmt.Println("Request handled by Sensor Server 2")
		return serverport2
	} else {
		fmt.Println("Request handled by Sensor Server 1")
		return serverport1
	}
}
