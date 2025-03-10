package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"servicePayment/src/payments/application"
	"servicePayment/src/payments/domain"
	"net/http"

	app "servicePayment/src/services/rabbitmq/application"
)

type CreatePaymentController struct {
    publishPaymentService *app.PublishPayment
	useCase *application.CreatePaymentUseCase
}

// NewCreatePaymentController crea un nuevo controlador para manejar pagos.
func NewCreatePaymentController(useCase *application.CreatePaymentUseCase, publishPaymentService *app.PublishPayment) *CreatePaymentController {
    return &CreatePaymentController{
        useCase:               useCase,
        publishPaymentService: publishPaymentService,
    }
}

// CreatePayment procesa una solicitud para crear un pago.
func (c *CreatePaymentController) CreatePayment(ctx *gin.Context) {
	var payment domain.Payment

	// 🔹 Log para verificar que la solicitud llega aquí
	log.Println("🟢 Se recibió una solicitud para crear un pago.")

	// Bind JSON to Payment struct
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		log.Println("❌ Error en datos recibidos:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// 🔹 Log para mostrar qué datos se están procesando
	log.Printf("📦 Datos del pago: %+v", payment)

	// Ejecutamos el caso de uso para guardar el pago.
	if err := c.useCase.Execute(&payment); err != nil {
		log.Println("❌ Error al procesar el pago:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// 🔹 Log antes de publicar en RabbitMQ
	log.Println("📤 Intentando enviar el pago a RabbitMQ...")

	// Publicar en RabbitMQ
	err := c.publishPaymentService.Execute(payment)
	if err != nil {
		log.Printf("❌ Error al publicar en RabbitMQ: %v", err)
	} else {
		log.Println("✅ Pago enviado correctamente a RabbitMQ.")
	}

	// Si el pago fue procesado correctamente, respondemos con éxito.
	ctx.JSON(http.StatusOK, gin.H{"status": "Payment processed successfully"})
}
