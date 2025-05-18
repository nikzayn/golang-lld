package models

import (
	"errors"
	"sync"
	"time"

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

func (s *BaseSeat) IsAvailable() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Status == enums.Available
}

func (s *BaseSeat) Book() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Status != enums.Available {
		return errors.New("seat already booked or reserved")
	}
	s.Status = enums.Booked
	return nil
}

func (s *BaseSeat) GetStatus() enums.SeatStatus {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Status
}

// Seat Types
type Platinum struct{ BaseSeat }
type Gold struct{ BaseSeat }
type Silver struct{ BaseSeat }

func (p *Platinum) SetSeat() {}
func (p *Platinum) SetRate() { p.Rate = 500 }

func (p *Gold) SetSeat()
func (p *Gold) SetRate() { p.Rate = 300 }

func (p *Silver) SetSeat()
func (p *Silver) SetRate() { p.Rate = 200 }

// Movie Schema
type Movie struct {
	Title                string
	Duration             int
	Genre                string
	ReleaseDate          time.Time
	Language             string
	IsSubtitlesAvailable bool
	Show                 []*ShowTime
}

// ShowTime Schema
type ShowTime struct {
	ShowID   int
	Start    time.Time
	Date     time.Time
	Duration int
	Seats    []Seat
}

func (s *ShowTime) ShowAvailableSeats() []Seat {
	available := []Seat{}
	for _, seat := range s.Seats {
		available = append(available, seat)
	}
	return available
}

// MovieTicket Schema
type MovieTicket struct {
	TicketID int
	Seat     Seat
	Movie    *Movie
	Show     *ShowTime
}

// City Schema
type City struct {
	Name    string
	State   string
	ZipCode string
	Cinemas []*Cinema
}

// Cinema Schema
type Cinema struct {
	CinemaID int
	Halls    []*Hall
	City     *City
}

type Hall struct {
	HallID int
	Shows  []*ShowTime
}
