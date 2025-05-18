package airlinemanagementsystem

import "sync"

type Payment struct {
	PaymentID     int
	PaymentMethod string
	Amount        float64
	status        PaymentStatus
	mu            sync.Mutex
}

func NewPayment(paymentId int, paymentMethod string, amount float64) *Payment {
	return &Payment{
		PaymentID:     paymentId,
		PaymentMethod: paymentMethod,
		Amount:        amount,
		status:        PaymentStatusPending,
	}
}

func (p *Payment) ProcessPayment() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.status = PaymentStatusCompleted
}
