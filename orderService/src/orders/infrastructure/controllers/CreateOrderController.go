package controllers

import (
	"log"
	"net/http"
	"serviceOrders/src/orders/application"
	"serviceOrders/src/orders/domain"
	app "serviceOrders/src/services/rabbitmq/application"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
    useCase             *application.CreateOrderUseCase
    publishOrderService *app.PublishOrder
}

// Constructor actualizado
func NewCreateOrderController(useCase *application.CreateOrderUseCase, publishOrderService *app.PublishOrder) *CreateOrderController {
    return &CreateOrderController{
        useCase:             useCase,
        publishOrderService: publishOrderService,
    }
}

func (c *CreateOrderController) CreateOrder(ctx *gin.Context) {
    var requestData struct {
        Name        string  `json:"name" binding:"required"`
        Amount      int     `json:"amount" binding:"required"`
        Description *string `json:"description,omitempty"` // Puede ser nil
    }

    if err := ctx.ShouldBindJSON(&requestData); err != nil {
        log.Printf("Error al recibir los datos: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
        return
    }

    // Crear la orden correctamente
    order := domain.NewOrder(requestData.Name, requestData.Amount, requestData.Description)

    // Llamamos al caso de uso para procesar la orden en la base de datos
    err := c.useCase.Execute(order.Name, order.Amount, &order.Description)
    if err != nil {
        log.Printf("Error al procesar la orden: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process order"})
        return
    }

    // Enviar a RabbitMQ con logs detallados
    log.Printf("Enviando mensaje a RabbitMQ: %+v", order)
    if err := c.publishOrderService.Execute(*order); err != nil {
        log.Printf("Error al enviar mensaje a RabbitMQ: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to RabbitMQ"})
        return
    }

    log.Printf("Mensaje enviado correctamente a RabbitMQ: %+v", order)
    ctx.JSON(http.StatusOK, gin.H{"status": "Order processed and message sent to RabbitMQ"})
}

