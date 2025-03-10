package routers

import (
	"github.com/gin-gonic/gin"
	"servicePayment/src/payments/infrastructure/controllers"
)

// PaymentRouter define las rutas relacionadas con los pagos.
func PaymentRouter(router *gin.Engine, controller *controllers.CreatePaymentController) {
	// Ruta para crear un nuevo pago.
	router.POST("/payments", controller.CreatePayment)
}
