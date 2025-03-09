
package application

import (
    "serviceOrders/src/orders/domain"
)

type CreatePaymentUseCase struct {
    paymentRepo domain.PaymentRepository
}

func NewCreatePaymentUseCase(paymentRepo domain.PaymentRepository) *CreatePaymentUseCase {
    return &CreatePaymentUseCase{paymentRepo: paymentRepo}
}

func (uc *CreatePaymentUseCase) Execute(name string, amount int, description string) (*domain.Payment, error) {
    payment := domain.NewPayment(name, amount, description)
    err := uc.paymentRepo.Save(payment)
    if err != nil {
        return nil, err
    }
    return payment, nil
}
