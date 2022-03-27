package structural

import "fmt"

type Seller interface {
	CancelReservation(charge float64)
}
type Reservation struct {
	sellerRef Seller // this is the implementer reference
}

func (r Reservation) Cancel() {
	r.sellerRef.CancelReservation(10) // charge $10 as cancellation feed
}

type PremiumReservation struct {
	Reservation
}

func (r PremiumReservation) Cancel() {
	r.sellerRef.CancelReservation(0) // no charges
}

type InstitutionSeller struct{}

func (s InstitutionSeller) CancelReservation(charge float64) {
	fmt.Println("InstitutionSeller CancelReservation charge =", charge)
}

type SmallScaleSeller struct{}

func (s SmallScaleSeller) CancelReservation(charge float64) {
	fmt.Println("SmallScaleSeller CancelReservation charge =", charge)
}
