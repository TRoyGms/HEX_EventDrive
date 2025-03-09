package main

import (
    "log"
    "serviceOrders/src/database"
    "serviceOrders/src/orders/application"
    "serviceOrders/src/orders/infrastructure/controllers"
    "serviceOrders/src/orders/persistence"
    "serviceOrders/src/services/rabbitmq/infraestructure"
    app "serviceOrders/src/services/rabbitmq/application" // Cambié a la importación correcta

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    // Initialize database and dependencies
    db, err := database.NewDBConnection()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Create RabbitMQ connection and service
    rabbitMQ, err := infraestructure.NewRabbitMQBus("amqp://guest:guest@54.175.122.102:5672/")
    if err != nil {
        log.Fatal("Failed to connect to RabbitMQ:", err)
    }
    defer rabbitMQ.Close()

    // Create services and handlers
    paymentRepo := persistence.NewMysqlPaymentRepository(db)
    createPaymentUseCase := application.NewCreatePaymentUseCase(paymentRepo)
    createPaymentController := controllers.NewCreatePaymentController(createPaymentUseCase)

    // Create PublishOrder service, passing RabbitMQ as the message bus
    publishOrderService := app.NewPublishOrder(rabbitMQ) // Usamos la importación correcta

    // Initialize Gin router
    router := gin.Default()

    // Define routes
    router.POST("/orders", func(c *gin.Context) {
        // Call the controller to create the order
        createPaymentController.CreateOrder(c)

        // After creating the order, send a message to RabbitMQ queue
        msg := []byte("New order created!")
        err := publishOrderService.Execute("orderQueue", msg)
        if err != nil {
            log.Println("Failed to publish message to RabbitMQ:", err)
        } else {
            log.Println("Message sent to RabbitMQ: New order created!")
        }
    })

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
