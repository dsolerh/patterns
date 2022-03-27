package creational

import "github.com/dsolerh/patterns/common"

func NewReservation(vertical, reservationDate string) common.Reservation {
	switch vertical {
	case "flight":
		return common.FlightReservation{ReservationDate: reservationDate}
	case "hotel":
		return common.HotelReservation{ReservationDate: reservationDate}
	default:
		return nil
	}
}
