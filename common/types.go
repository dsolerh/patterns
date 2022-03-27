package common

type Reservation interface{}
type Invoice interface{}

type FlightReservation struct {
	ReservationDate string
}

type FlightInvoice struct{}

type HotelReservation struct {
	ReservationDate string
}

type HotelInvoice struct{}
