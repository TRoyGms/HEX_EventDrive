package infrastructure

import (
	"log"
	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

func NewWebSocketClient(url string) (*WebSocketClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocketClient{conn: conn}, nil
}

func (ws *WebSocketClient) NotifySocket(message string) error {
	if ws.conn == nil {
		return nil
	}
	
	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error sending WebSocket message:", err)
		return err
	}
	return nil
}

func (ws *WebSocketClient) Close() {
	if ws.conn != nil {
		ws.conn.Close()
	}
}
