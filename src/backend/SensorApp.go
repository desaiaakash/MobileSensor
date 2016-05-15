package main

import "backend/sensor"
import "os"
import "strconv"

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	sensor.StartServer(port)
}
