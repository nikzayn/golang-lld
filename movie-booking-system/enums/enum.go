package enums

type PaymentStatus int
type BookingStatus int
type SeatStatus int

const (
	Pending PaymentStatus = iota
	Confirmed
	Declined
	Refunded
)

const (
	BookingPending BookingStatus = iota
	BookingConfirmed
	Cancelled
	Denied
	BookingRefund
)

const (
	Available SeatStatus = iota
	Booked
	Reserved
)
