package infrastructure

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
	upgrader websocket.Upgrader
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar la conexión:", err)
		return
	}

	ws.mu.Lock()
	ws.clients[conn] = true
	ws.mu.Unlock()

	log.Println("Cliente conectado. Total de clientes:", len(ws.clients)) // 🔥 Agregar este log

	// Mantener la conexión abierta
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			ws.mu.Lock()
			delete(ws.clients, conn)
			ws.mu.Unlock()
			conn.Close()
			break
		}
	}
}



func (ws *WebSocketServer) BroadcastMessage(message string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	log.Println("Intentando enviar mensaje a WebSocket:", message) // 🔥 Agrega esta línea

	for client := range ws.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error enviando mensaje al WebSocket:", err)
			client.Close()
			delete(ws.clients, client)
		} else {
			log.Println("Mensaje enviado a un cliente WebSocket")
		}
	}
}



func (ws *WebSocketServer) NotifySocket(message string) error {
	ws.BroadcastMessage(message)
	return nil
}

func (ws *WebSocketServer) Start(port string) {
	http.HandleFunc("/ws", ws.HandleConnections)
	log.Println("WebSocket server started on port", port)
	go http.ListenAndServe(port, nil)
}

