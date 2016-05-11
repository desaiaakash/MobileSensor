package sensor

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

//function to get current sensor value

func getSensorValue(sensorId string) SensorValueRes {

	sdres := SensorDataRes{}
	svres := SensorValueRes{}

	//connect to mongo db
	session, err1 := mgo.Dial(getDatabaseURL())

	// Check if connection error, is mongo running?
	if err1 != nil {
		fmt.Println("Error while connecting to database")
		os.Exit(1)
	}

	err2 := session.DB("sensordb").C("sensor").FindId(sensorId).One(&sdres)

	if err2 != nil {
		panic("Error while fetching document from DB " + err2.Error())
	}

	//update value and history
	sdres.History = append(sdres.History, sdres.Value)
	sdres.Value = getSensorData(sdres.Coordinate.Lat, sdres.Coordinate.Lng, sdres.TypeCode)

	//update document in mongoDB
	err3 := session.DB("sensordb").C("sensor").UpdateId(sensorId, sdres)

	if err3 != nil {
		panic(err3)
	}

	session.Close()

	//prepare response
	svres.Value = sdres.Value
	svres.History = sdres.History

	return svres
}

// function which removes sensor from database
func removeSensor(sensorId string) {

	//connect to mongo db
	session, err1 := mgo.Dial(getDatabaseURL())

	// Check if connection error, is mongo running?
	if err1 != nil {
		fmt.Println("Error while connecting to database")
		os.Exit(1)
	}

	// Remove sensor
	err2 := session.DB("sensordb").C("sensor").RemoveId(sensorId)
	if err2 != nil {
		panic("Panic while removing a sensor " + err2.Error())
	}

	session.Close()

}

// function which adds sensor in the database
func addSensor(data AddSensorReq) SensorDataRes {

	// empty struct for returning data
	sdres := SensorDataRes{}

	//copying corresponding values
	sdres.Type = data.Type
	sdres.Location = data.Location
	sdres.Coordinate.Lat = data.Coordinate.Lat
	sdres.Coordinate.Lng = data.Coordinate.Lng
	sdres.TypeCode = data.TypeCode

	//set Color bases on type code
	sdres.Color = getSensorColor(data.TypeCode)

	//set status - since the sensor is just created, its status will be activated
	sdres.Status = true

	//fetch value corresponding to a sensor type
	sdres.Value = getSensorData(data.Coordinate.Lat, data.Coordinate.Lng, data.TypeCode)

	//get cost - will be random based on type
	sdres.Cost = getSensorCost(data.TypeCode)

	//get id
	sdres.ID = bson.NewObjectId().Hex()

	//set history
	historySlice := make([]float64, 0)
	sdres.History = historySlice

	//connect to mongo db
	session, err := mgo.Dial(getDatabaseURL())

	// Check if connection error, is mongo running?
	if err != nil {
		panic("Error while connecting to database")
		os.Exit(1)
	}

	//add the sensor to sensor collection in sensordb database
	session.DB("sensordb").C("sensor").Insert(sdres)

	//close the session
	session.Close()

	return sdres
}

//activate sensor - change sensor status
func changeSensorStatus(sensorId string, status bool) SensorDataRes {

	// empty struct for returning data
	sdres := SensorDataRes{}

	//connect to mongo db
	session, err1 := mgo.Dial(getDatabaseURL())

	// Check if connection error, is mongo running?
	if err1 != nil {
		fmt.Println("Error while connecting to database")
		os.Exit(1)
	}

	err2 := session.DB("sensordb").C("sensor").FindId(sensorId).One(&sdres)

	if err2 != nil {
		panic("Error while fetching document from DB " + err2.Error())
	}

	//update status
	sdres.Status = status

	//update document in mongoDB
	err3 := session.DB("sensordb").C("sensor").UpdateId(sensorId, sdres)

	if err3 != nil {
		panic(err3)
	}
	session.Close()

	return sdres
}

func getSensorList(data SensorListReq) SensorListRes {

	slres := SensorListRes{}

	//connect to mongo db
	session, err1 := mgo.Dial(getDatabaseURL())

	// Check if connection error, is mongo running?
	if err1 != nil {
		fmt.Println("Error while connecting to database")
		os.Exit(1)
	}

	//get the slice from the struct which contain the all the sensorIds
	sensorIdList := data.List

	for _, sid := range sensorIdList {

		// empty struct for returning data
		sdres := SensorDataRes{}

		//get data for particular sensor id
		err2 := session.DB("sensordb").C("sensor").FindId(sid).One(&sdres)

		if err2 != nil {
			panic("Error while fetching document from DB " + err2.Error())
		}

		//update each sensor value by fetching latest value
		sdres.History = append(sdres.History, sdres.Value)
		sdres.Value = getSensorData(sdres.Coordinate.Lat, sdres.Coordinate.Lng, sdres.TypeCode)

		//update document in mongoDBx
		err3 := session.DB("sensordb").C("sensor").UpdateId(sid, sdres)

		if err3 != nil {
			panic(err3)
		}

		//add the structure to list
		slres.Result = append(slres.Result, sdres)
	}

	session.Close()

	return slres
}

func getDatabaseURL() string {
	var dburl string = "mongodb://shikher9:mongo1239@ds033103.mlab.com:33103/sensordb"
	return dburl
}
