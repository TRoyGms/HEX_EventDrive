package domain

type PaymentRepository interface {
	Save(payment *Payment) error
}
