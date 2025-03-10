package main

import (
	"log"
	"serviceOrders/src/database"
	"serviceOrders/src/orders/application"
	"serviceOrders/src/orders/infrastructure/controllers"
	"serviceOrders/src/orders/infrastructure/interfaces/http/routers"
	"serviceOrders/src/orders/persistence"
	app "serviceOrders/src/services/rabbitmq/application"
	"serviceOrders/src/services/rabbitmq/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	rabbitMQ, err := infraestructure.NewRabbitMQBus("amqp://guest:guest@54.175.122.102:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rabbitMQ.Close()

	orderRepo := persistence.NewMysqlOrderRepository(db)
	createOrderUseCase := application.NewCreateOrderUseCase(orderRepo)
	publishOrderService := app.NewPublishOrder(rabbitMQ)

	createOrderController := controllers.NewCreateOrderController(createOrderUseCase, publishOrderService)

	router := gin.Default()

	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes (¡Cuidado en producción!)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routers.RegisterOrderRoutes(router, createOrderController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
