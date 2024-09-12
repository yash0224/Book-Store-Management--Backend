package models

import (
	"fmt"
	"log"
)

// var dbs *sql.DB

type Orders struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact string `json:"contact"`
	Book    int    `json:"bookId"`
}

// fetch the orders for particular book
func BookOrders(bookId int) ([]Orders, error) {
	query := "SELECT id, name, address, contact, bookId FROM Orders WHERE bookId = ?"

	rows, err := db.Query(query, bookId)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var orders []Orders

	for rows.Next() {
		var order Orders
		if err := rows.Scan(&order.ID, &order.Name, &order.Address, &order.Contact, &order.Book); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return orders, nil
}

// fetch all orders
func GetAllOrders() ([]Orders, error) {
	query := "SELECT id, name, address, contact, bookId FROM Orders"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var orders []Orders

	for rows.Next() {
		var order Orders
		if err := rows.Scan(&order.ID, &order.Name, &order.Address, &order.Contact, &order.Book); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return orders, nil
}

// insert a new order
func InsertOrder(o *Orders) (*Orders, error) {
	query := `
		INSERT INTO Orders (name, address, contact, bookId)
		VALUES (?, ?, ?, ?)
	`
	result, err := db.Exec(query, o.Name, o.Address, o.Contact, o.Book)
	if err != nil {
		return nil, fmt.Errorf("failed to insert order: %v", err)
	}

	if err := UpdateStock(o.Book); err != nil {
		log.Printf("Failed to update stock: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	o.ID = int(id)

	return o, nil
}
