package parking

import (
	"testing"
)

func TestIfDriverCanPark(t *testing.T) {
	vehicle, _ := AssignNewParkingSpace("aa00bb1111", "car", true)
	got := parkingSpace{Vehicle{"aa00bb1111", "car"}, Ticket{true, 1}}
	if *vehicle != got {
		t.Errorf("Error! Wrong entry")
	}
}

func TestIfParkingLotHasAnEmptySlotForANewCar(t *testing.T) {
	_, err := AssignNewParkingSpace("aa00bb1111", "car", true)
	if err != nil {
		t.Errorf("Error! Parking lot has no empty slots. ")
	}
}
func TestIfCarWasUnparkedFromTheLot(t *testing.T) {
	err := UnparkAnAlreadyParkedCar(parkingSpace{Vehicle{"aa00bb1111", "car"}, Ticket{true, 1}})
	if err == VehicleNotParkedError {
		t.Errorf("Error not parked")
	}

}
