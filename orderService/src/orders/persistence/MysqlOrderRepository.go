package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"serviceOrders/src/orders/domain"
)

// MysqlOrderRepository es un repositorio para interactuar con la base de datos MySQL.
type MysqlOrderRepository struct {
	db *sql.DB
}

// NewMysqlOrderRepository crea un nuevo repositorio de órdenes.
func NewMysqlOrderRepository(db *sql.DB) *MysqlOrderRepository {
	return &MysqlOrderRepository{db: db}
}

// Save guarda una nueva orden en la base de datos.
func (r *MysqlOrderRepository) Save(order *domain.Order) error {
	// Preparamos la consulta SQL para insertar la orden en la base de datos
	query := `INSERT INTO orders (name, amount, description) VALUES (?, ?, ?)`

	// Ejecutamos la consulta con los datos de la orden
	_, err := r.db.Exec(query, order.Name, order.Amount, order.Description)
	if err != nil {
		log.Printf("Error saving order to database: %v", err) // Log para ver errores
		return fmt.Errorf("error al guardar la orden: %v", err)
	}

	// Si no hay errores, la orden fue guardada con éxito
	return nil
}

