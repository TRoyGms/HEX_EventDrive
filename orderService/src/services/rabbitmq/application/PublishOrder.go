package application

import (
	"encoding/json"
	"errors"
	"log"
	"serviceOrders/src/orders/domain"
	"serviceOrders/src/services/rabbitmq/domain/repositories"
)

type PublishOrder struct {
	messageBus repositories.IMessageBus
}

func NewPublishOrder(messageBus repositories.IMessageBus) *PublishOrder {
	return &PublishOrder{
		messageBus: messageBus,
	}
}

// Execute maneja la lógica de la publicación del mensaje a RabbitMQ.
func (s *PublishOrder) Execute(order domain.Order) error {
	if order.Name == "" || order.Amount <= 0 {
		log.Println("Error: La orden no tiene un nombre o el monto es inválido")
		return errors.New("order data is invalid")
	}

	// Convertir la estructura `order` a JSON
	msgBytes, err := json.Marshal(order)
	if err != nil {
		log.Printf("Error al serializar la orden: %v", err)
		return err
	}

	// Log para verificar si el mensaje se está generando correctamente
	log.Printf("Enviando mensaje a RabbitMQ: %s", string(msgBytes))

	// Asegurar que la conexión al messageBus no sea nula
	if s.messageBus == nil {
		log.Println("Error: messageBus es nil, no se puede enviar mensaje a RabbitMQ")
		return errors.New("messageBus is nil")
	}

	// Enviar el mensaje a la cola RabbitMQ
	err = s.messageBus.Publish("orderQueue", msgBytes)
	if err != nil {
		log.Printf("Error al publicar el mensaje a RabbitMQ: %v", err)
		return err
	}

	log.Printf("Mensaje enviado correctamente a RabbitMQ: %+v", order)
	return nil
}
