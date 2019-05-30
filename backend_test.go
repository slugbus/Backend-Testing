package test

import (
	"fmt"
	"testing"
	"time"

	merge "github.com/slugbus/backend-merge"
	"github.com/slugbus/taps"
)

func TestTapsMergeMeasurements(*testing.T) {
	currentBuses := taps.UpdatedBusMap{}
	// get a ping of the buses running
	pingedBus, _ := taps.QueryAsMap()
	// print out pingedBus information
	fmt.Printf("First ping:\n")
	for k := range pingedBus {
		fmt.Printf("ID: %s\n Lon: %f\n Lat: %f\n Type: %s\n", pingedBus[k].ID, pingedBus[k].Lon, pingedBus[k].Lat, pingedBus[k].Type)
	}
	// merge with current bus state
	currentBuses = merge.MergeWithState(pingedBus, 0.0, currentBuses)
	// wait 3 seconds to ping taps server once again

	time.Sleep(3 * time.Second)
	//ping the server again
	pingedBus, _ = taps.QueryAsMap()
	//merge with currentBuses
	currentBuses = merge.MergeWithState(pingedBus, 3000.0, currentBuses)
	fmt.Printf("\nMerged Data:\n")
	// print out currentBuses
	for k := range currentBuses {
		fmt.Printf("ID: %s\n location: %f %f\n Type: %s\n Speed: %f\n Angle: %f\n", currentBuses[k].ID, currentBuses[k].Lat, currentBuses[k].Lon, currentBuses[k].Type, currentBuses[k].Speed, currentBuses[k].Angle)
	}
}
