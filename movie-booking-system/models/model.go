package models

import (
	"sync"

	"github.com/nikzayn/lld-golang/movie-booking-system/enums"
)

// Person Details
type Person struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// Customer schema
type Customer struct {
	Person
	Bookings []*Booking
}

// Customer operations
func (c *Customer) CreateBooking(b *Booking) bool {
	c.Bookings = append(c.Bookings, b)
	return true
}

func (c *Customer) ModifyBooking(b *Booking) bool {
	return true
}

func (c *Customer) DeleteBooking(b *Booking) bool {
	return true
}

// Admin Schema
type Admin struct {
	Person
}

func (a *Admin) AddShow(show *ShowTime) bool    { return true }
func (a *Admin) UpdateShow(show *ShowTime) bool { return true }
func (a *Admin) DeleteShow(show *ShowTime) bool { return true }
func (a *Admin) AddMovie(movie *Movie) bool     { return true }
func (a *Admin) DeleteMovie(movie *Movie) bool  { return true }

// Ticket Schema
type Ticket struct {
	Person
}

func (t *Ticket) CreateBooking(b *Booking) bool { return true }

type Seat interface {
	IsAvailable() bool
	SetSeat()
	SetRate()
	Book() error
	GetStatus() enums.SeatStatus
}

type BaseSeat struct {
	SeatNo string
	Status enums.SeatStatus
	Rate   float64
	mu     sync.Mutex
}
