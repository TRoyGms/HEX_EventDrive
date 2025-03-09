package app

import "serviceOrders/src/services/rabbitmq/domain/repositories"

type PublishOrder struct {
	messageBus repositories.IMessageBus
}

func NewPublishOrder(messageBus repositories.IMessageBus) *PublishOrder {
	return &PublishOrder{
		messageBus: messageBus,
	}
}

func (s *PublishOrder) Execute(queue string, msg []byte) error {
	return s.messageBus.Publish(queue, msg)
}
