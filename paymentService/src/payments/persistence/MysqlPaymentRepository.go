package persistence

import (
	"servicePayment/src/payments/domain"
	"database/sql"
	"fmt"
)

// MysqlPaymentRepository interactúa con la base de datos MySQL.
type MysqlPaymentRepository struct {
	db *sql.DB
}

// NewMysqlPaymentRepository crea una nueva instancia del repositorio.
func NewMysqlPaymentRepository(db *sql.DB) *MysqlPaymentRepository {
	return &MysqlPaymentRepository{db: db}
}

// Save guarda un nuevo pago en la base de datos.
func (r *MysqlPaymentRepository) Save(payment *domain.Payment) error {
	// Preparamos la consulta SQL para insertar el pago.
	query := `INSERT INTO payments (name, amount, description) VALUES (?, ?, ?)`

	// Ejecutamos la consulta con los datos del pago.
	_, err := r.db.Exec(query, payment.Name, payment.Amount, payment.Description)
	if err != nil {
		return fmt.Errorf("error al guardar el pago: %v", err)
	}

	return nil
}
