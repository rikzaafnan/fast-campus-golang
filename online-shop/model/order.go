package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Checkout struct {
	Email    string            `json:"email"`
	Address  string            `json:"address"`
	Products []ProductQuantity `json:"products"`
}

type ProductQuantity struct {
	ID       string `json:"id"`
	Quantity int32  `json:"quantity"`
}

type Order struct {
	ID                string     `json:"id"`
	Email             string     `json:"email"`
	Address           string     `json:"address"`
	GrandTotal        int64      `json:"grand_total"`
	Passcode          *string    `json:"passcode,omitempty"`
	PaidAt            *time.Time `json:"paid_at,omitempty"`
	PaidBank          *string    `json:"paid_bank,omitempty"`
	PaidAccountNumber *string    `json:"paid_account_number,omitempty"`
}

type OrderDetail struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
	Price     int64  `json:"price"`
	Total     int64  `json:"total"`
}

type OrderWithDetail struct {
	Order
	Detail []OrderDetail `json:"detail"`
}

type Confirm struct {
	Amount        int64  `json:"amount" binding:"required"`
	Bank          string `json:"bank" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	Passcode      string `json:"passcode" binding:"required"`
}

func CreateOrder(db *sql.DB, order Order, details []OrderDetail) error {
	if db == nil {
		return errDBNil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	queryOrder := `INSERT INTO orders (id, email, address, passcode, grand_total) VALUES ($1, $2,$3,$4,$5)`
	_, err = tx.Exec(queryOrder, order.ID, order.Email, order.Address, order.Passcode, order.GrandTotal)
	if err != nil {
		tx.Rollback()
		return err
	}

	//queryOrderDetail := `INSERT INTO order_details (id, order_id, product_id, quantity, price, total) VALUES ($1, $2, $3, $4, $5,$6)`
	//for _, d := range details {
	//	_, err = tx.Exec(queryOrderDetail, d.ID, d.OrderID, d.ProductID, d.Quantity, d.Price, d.Total)
	//	if err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//}

	//TODO : improve code
	queryOrderDetail := `INSERT INTO order_details (id, order_id, product_id, quantity, price, total) VALUES `

	var valueStrings []string
	var valueArgs []interface{}

	for i, d := range details {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
		valueArgs = append(valueArgs, d.ID, d.OrderID, d.ProductID, d.Quantity, d.Price, d.Total)
	}

	queryOrderDetail += strings.Join(valueStrings, ",")

	_, err = tx.Exec(queryOrderDetail, valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func SelectOrderByID(db *sql.DB, id string) (Order, error) {
	if db == nil {
		return Order{}, errDBNil
	}

	query := `SELECT id, email, address, passcode, grand_total, paid_at, paid_bank, paid_account FROM orders WHERE id = $1`
	var order Order

	row := db.QueryRow(query, id)

	err := row.Scan(&order.ID, &order.Email, &order.Address, &order.Passcode, &order.GrandTotal, &order.PaidAt, &order.PaidBank, &order.PaidAccountNumber)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func UpdateOrderByID(db *sql.DB, id string, confirmReq Confirm, currentTime time.Time) error {
	if db == nil {
		return errDBNil
	}

	query := `UPDATE orders SET paid_at=$1, paid_bank=$2, paid_account=$3 WHERE id=$4`

	_, err := db.Exec(query, currentTime, confirmReq.Bank, confirmReq.AccountNumber, id)
	if err != nil {
		return err
	}

	return nil
}
