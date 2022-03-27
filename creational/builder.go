package creational

import "github.com/dsolerh/patterns/common"

type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() common.Reservation
}

type reservationBuilder struct {
	vertical string
	rdate    string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date
	return r
}

func (r *reservationBuilder) Build() common.Reservation {
	var builtReservation common.Reservation
	switch r.vertical {
	case "flight":
		builtReservation = common.FlightReservation{ReservationDate: r.rdate}
	case "hotel":
		builtReservation = common.HotelReservation{ReservationDate: r.rdate}
	}
	return builtReservation
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}
