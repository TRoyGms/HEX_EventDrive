package routers

import (
    "serviceOrders/src/orders/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

// RegisterOrderRoutes configura las rutas relacionadas con órdenes
func RegisterOrderRoutes(router *gin.Engine, createOrderController *controllers.CreateOrderController) {
    router.POST("/orders", createOrderController.CreateOrder)
}
