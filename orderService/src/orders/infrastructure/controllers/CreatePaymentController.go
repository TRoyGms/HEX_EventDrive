
package controllers

import (
    "net/http"
    "serviceOrders/src/orders/application"
    "github.com/gin-gonic/gin"
)

type CreatePaymentController struct {
    useCase *application.CreatePaymentUseCase
}

func NewCreatePaymentController(useCase *application.CreatePaymentUseCase) *CreatePaymentController {
    return &CreatePaymentController{useCase: useCase}
}

func (c *CreatePaymentController) CreateOrder(ctx *gin.Context) {
    var request struct {
        Name        string `json:"name"`
        Amount      int    `json:"amount"`
        Description string `json:"description"`
    }
    if err := ctx.BindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    payment, err := c.useCase.Execute(request.Name, request.Amount, request.Description)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"payment": payment})
}
