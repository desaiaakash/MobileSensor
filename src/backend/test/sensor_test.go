package test

import "testing"
import "backend/sensor"

func TestServerStart(t *testing.T) {
	sensor.StartServer(80)
}
