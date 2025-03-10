package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQBus(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Error al abrir un canal en RabbitMQ: %v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"notificationQueue",
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("❌ Error al declarar la cola notificationQueue: %v", err)
	}

	log.Println("✅ Conectado a RabbitMQ y cola notificationQueue declarada correctamente.")

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) Publish(queue string, msg []byte) error {
	log.Printf("📤 Intentando publicar mensaje en la cola %s", queue)

	log.Printf("📦 Contenido del mensaje: %s", string(msg)) // 🔹 Muestra el contenido del mensaje antes de enviarlo

	err := r.ch.Publish(
		"",    // Default exchange
		queue, // Queue name
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)

	if err != nil {
		log.Printf("❌ Error al publicar en RabbitMQ: %v", err)
	} else {
		log.Println("✅ Mensaje enviado a RabbitMQ correctamente.")
	}

	return err
}

func (r *RabbitMQ) Close() {
	log.Println("🔻 Cerrando conexión a RabbitMQ...")
	r.conn.Close()
	r.ch.Close()
	log.Println("🔻 Conexión cerrada correctamente.")
}
