package application

import "servicePayment/src/infrastructure"



type SocketUseCase struct {
	webSocketClient *infrastructure.WebSocketClient
}

func NewSocketUseCase(wsClient *infrastructure.WebSocketClient) *SocketUseCase {
	return &SocketUseCase{webSocketClient: wsClient}
}

func (s *SocketUseCase) NotifySocket(message string) error {
	return s.webSocketClient.NotifySocket(message)
}
