package application

import (
	"encoding/json"
	"errors"
	"log"
	"servicePayment/src/payments/domain"
	"servicePayment/src/services/rabbitmq/domain/repositories"
)

type PublishPayment struct {
	messageBus repositories.IMessageBus
}

func NewPublishPayment(messageBus repositories.IMessageBus) *PublishPayment {
	return &PublishPayment{
		messageBus: messageBus,
	}
}


func (s *PublishPayment) Execute(payment domain.Payment) error {
	if payment.Amount <= 0 {
		log.Println("Error: El monto es inválido")
		return errors.New("payment data is invalid")
	}

	
	msgBytes, err := json.Marshal(payment)
	if err != nil {
		log.Printf("Error al serializar el pago: %v", err)
		return err
	}


	log.Printf("Enviando mensaje a RabbitMQ: %s", string(msgBytes))
	return s.messageBus.Publish("notificationQueue", msgBytes)
}

