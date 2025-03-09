package persistence

import (
	"serviceOrders/src/orders/domain"
	"database/sql"
	"fmt"
)

// MysqlPaymentRepository es un repositorio para interactuar con la base de datos MySQL.
type MysqlPaymentRepository struct {
	db *sql.DB
}

// NewMysqlPaymentRepository crea un nuevo repositorio de pagos.
func NewMysqlPaymentRepository(db *sql.DB) *MysqlPaymentRepository {
	return &MysqlPaymentRepository{db: db}
}

// Save guarda un nuevo pago en la base de datos.
func (r *MysqlPaymentRepository) Save(payment *domain.Payment) error {
	// Preparamos la consulta SQL para insertar la orden en la base de datos
	query := `INSERT INTO payments (name, amount, description) VALUES (?, ?, ?)`

	// Ejecutamos la consulta con los datos del pago
	_, err := r.db.Exec(query, payment.Name, payment.Amount, payment.Description)
	if err != nil {
		return fmt.Errorf("error al guardar el pago: %v", err)
	}

	// Si no hay errores, el pago fue guardado con éxito
	return nil
}
