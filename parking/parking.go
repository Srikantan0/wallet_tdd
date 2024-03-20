package parking

import (
	"errors"
)

var NoFreeParkingSpaceError = errors.New("Parking Lot has no empty slots")
var VehicleDuplicateError = errors.New("Vehicle is Already parked. ")
var VehicleNotParkedError = errors.New("Vehicle is not parked to be unparked. ")

type Vehicle struct {
	vehicleNum  string
	vehicleType string
}

type Ticket struct {
	ifDriverHasTicket bool
	ticketID          int
}

type parkingSpace struct {
	vehicle Vehicle
	ticket  Ticket
}

func AssignNewParkingSpace(vehicleNum, vehicleType string, ticket bool) (*parkingSpace, error) {
	for ticketId, free := range parkingLot {
		if free {
			_, ok := parked[vehicleNum]
			if ok == true {
				parkingLot[ticketId] = false
				parked[vehicleNum] = true
				return &parkingSpace{Vehicle{vehicleNum, vehicleType}, Ticket{ticket, ticketId}}, nil
			} else {
				return nil, VehicleDuplicateError
			}
		}
	}
	return nil, NoFreeParkingSpaceError
}

func UnparkAnAlreadyParkedCar(p parkingSpace) error {
	_, parkErr := parked[p.vehicle.vehicleNum]
	if parkErr == false {
		return VehicleNotParkedError
	}
	parkingLot[p.ticket.ticketID] = true
	parked[p.vehicle.vehicleNum] = false
	return nil

}
