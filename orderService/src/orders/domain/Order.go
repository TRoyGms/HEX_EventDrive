package domain

// Order representa la estructura de la orden.
type Order struct {
    ID          int    `json:"id,omitempty"`               // ID opcional
    Name        string `json:"name" binding:"required"`   // Nombre obligatorio
    Amount      int    `json:"amount" binding:"required"` // Monto obligatorio
    Description string `json:"description,omitempty"`     // Descripción opcional
}

// Constructor sin ID, la base de datos lo asignará
func NewOrder(name string, amount int, description *string) *Order {
    desc := ""
    if description != nil {
        desc = *description
    }

    return &Order{
        Name:        name,
        Amount:      amount,
        Description: desc,
    }
}
