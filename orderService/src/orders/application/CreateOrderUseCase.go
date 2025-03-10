package application

import "serviceOrders/src/orders/domain"

// CreateOrderUseCase maneja la lógica para crear una orden
type CreateOrderUseCase struct {
	repo domain.OrderRepository
}

// NewCreateOrderUseCase crea una nueva instancia del caso de uso para crear una orden
func NewCreateOrderUseCase(repo domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{repo: repo}
}

// Execute ejecuta la lógica de creación de la orden
func (uc *CreateOrderUseCase) Execute(name string, amount int, description *string) error {
    order := domain.NewOrder(name, amount, description)
    return uc.repo.Save(order)
}

