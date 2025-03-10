
package domain

type OrderRepository interface {
    Save(order *Order) error
}
