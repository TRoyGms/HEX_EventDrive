package repositories

type IMessageBus interface {
	Publish(queue string, msg []byte) error
}