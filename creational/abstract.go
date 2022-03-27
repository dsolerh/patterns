package creational

import "github.com/dsolerh/patterns/common"

type AbstractFactory interface {
	CreateReservation() common.Reservation
	CreateInvoice() common.Invoice
}

type HotelFactory struct{}

func (f HotelFactory) CreateReservation() common.Reservation {
	return new(common.HotelReservation)
}
func (f HotelFactory) CreateInvoice() common.Invoice {
	return new(common.HotelInvoice)
}

type FlightFactory struct{}

func (f FlightFactory) CreateReservation() common.Reservation {
	return new(common.FlightReservation)
}
func (f FlightFactory) CreateInvoice() common.Invoice {
	return new(common.FlightReservation)
}

func GetFactory(vertical string) AbstractFactory {
	var factory AbstractFactory
	switch vertical {
	case "flight":
		factory = FlightFactory{}
	case "hotel":
		factory = HotelFactory{}
	}
	return factory
}
