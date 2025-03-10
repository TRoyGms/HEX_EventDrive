package repositories

// IMessageBus es la interfaz que define cómo interactuar con RabbitMQ
type IMessageBus interface {
	Publish(queue string, msg []byte) error
}
