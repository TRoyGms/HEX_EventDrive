package application

import (
	"encoding/json"
	"log"
	"servicePayment/src/payments/domain"
)

type PaymentMessage struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type CreatePaymentUseCase struct {
	repository    domain.PaymentRepository
	socketUseCase ISocketUseCase
}

func NewCreatePaymentUseCase(repo domain.PaymentRepository, socket ISocketUseCase) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{repository: repo, socketUseCase: socket}
}

func (c *CreatePaymentUseCase) Execute(payment *domain.Payment) error {
	// Guardar el pago en la base de datos
	err := c.repository.Save(payment)
	if err != nil {
		return err
	}

	// Crear un objeto con la información del pago
	messageData := PaymentMessage{
		Name:   payment.Name,
		Amount: payment.Amount,
	}

	// Convertir a JSON
	messageJSON, err := json.Marshal(messageData)
	if err != nil {
		log.Println("Error al convertir el mensaje a JSON:", err)
		return err
	}

	log.Println("Enviando mensaje al WebSocket:", string(messageJSON))

	// Enviar mensaje al WebSocket
	c.socketUseCase.NotifySocket(string(messageJSON))

	return nil
}
