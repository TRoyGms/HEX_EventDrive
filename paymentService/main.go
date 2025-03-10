package main

import (
	"log"
	"servicePayment/src/payments/application"
	"servicePayment/src/payments/infrastructure/controllers"
	"servicePayment/src/payments/persistence"
	services "servicePayment/src/services/rabbitmq/infraestructure"
	"servicePayment/src/payments/infrastructure/interfaces/http/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"servicePayment/src/database"
	websocket "servicePayment/src/infrastructure"

	rabbitmqApp "servicePayment/src/services/rabbitmq/application"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Inicializar RabbitMQ (opcional)
	rmq, err := services.NewRabbitMQBus("amqp://guest:guest@54.175.122.102:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rmq.Close()

	// Inicializar la conexión a la base de datos MySQL
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Inicializar el servidor WebSocket
	wsServer := websocket.NewWebSocketServer()
	go wsServer.Start(":8000") // WebSocket en el puerto 8081

	// Crear el repositorio y el caso de uso con WebSocket integrado
	repo := persistence.NewMysqlPaymentRepository(db)
	processPayment := application.NewCreatePaymentUseCase(repo, wsServer)

	// Crear el servicio de publicación en RabbitMQ
	publishPaymentService := rabbitmqApp.NewPublishPayment(rmq)

	// Crear el controlador
	paymentHandler := controllers.NewCreatePaymentController(processPayment, publishPaymentService)

	// Configurar Gin (servidor HTTP)
	router := gin.Default()

	// Enrutar las solicitudes para procesar los pagos
	routers.PaymentRouter(router, paymentHandler)

	// Iniciar el servidor API en el puerto 8082
	if err := router.Run(":8084"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
