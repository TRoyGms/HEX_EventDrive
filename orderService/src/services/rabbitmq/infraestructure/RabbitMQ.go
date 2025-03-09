package infraestructure

import "github.com/streadway/amqp"

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (r *RabbitMQ) Close() {
	panic("unimplemented")
}

func NewRabbitMQBus(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Declarar la cola 'orderQueue' para asegurar que existe y es duradera
	_, err = ch.QueueDeclare(
		"orderQueue", // Nombre de la cola
		true,         // Durable (se mantiene incluso si RabbitMQ se reinicia)
		false,        // Auto-delete (no se elimina automáticamente)
		false,        // Exclusive (puede ser usada por otras conexiones)
		false,        // No-wait (sin esperar una respuesta)
		nil,          // Argumentos adicionales
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Publish(queue string, msg []byte) error {
	return r.ch.Publish(
		"",    // Exchange vacío, lo que significa que la cola se usará directamente
		queue, // Nombre de la cola
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}
