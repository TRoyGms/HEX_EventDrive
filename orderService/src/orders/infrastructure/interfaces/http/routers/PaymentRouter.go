
package routers

import (
    "serviceOrders/src/orders/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.Engine, controller *controllers.CreatePaymentController) {
    router.POST("/orders", controller.CreateOrder)
}
