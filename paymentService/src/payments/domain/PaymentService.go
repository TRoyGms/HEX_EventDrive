package domain

import (
	"fmt"
)

// PaymentService contiene la lógica de negocio para procesar pagos
type PaymentService struct {
	repo PaymentRepository
}

// NewPaymentService crea una nueva instancia del servicio de pagos
func NewPaymentService(repo PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

// ProcessPayment procesa una orden de pago y guarda la información del pago
func (ps *PaymentService) ProcessPayment(payment *Payment) error {
	// Aquí solo almacenamos el pago, ya que la validación se hará en el frontend
	err := ps.repo.Save(payment)
	if err != nil {
		return fmt.Errorf("error al guardar el pago: %v", err)
	}

	// Si todo sale bien, devolvemos nil (sin error)
	return nil
}
