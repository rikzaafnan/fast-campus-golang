package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	ID        string `json:"id" binding:"len=0"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	IsDeleted *bool  `json:"is_deleted,omitempty"`
}

var (
	errDBNil = errors.New("connection is not establish")
)

func SelectProduct(db *sql.DB) ([]Product, error) {

	if db == nil {
		return nil, errDBNil
	}

	query := ` SELECT id, name, price FROM products WHERE is_deleted = false `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []Product{}

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil

}

func SelectProductByID(db *sql.DB, id string) (Product, error) {
	if db == nil {
		return Product{}, errDBNil
	}

	query := `SELECT id, name , price FROM products WHERE is_deleted = false AND id = $1`

	var product Product
	row := db.QueryRow(query, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return Product{}, err
	}

	return product, nil

}

func InsertProduct(db *sql.DB, product Product) error {
	if db == nil {
		return errDBNil
	}

	query := `INSERT INTO products (id, name, price) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(db *sql.DB, product Product) error {
	if db == nil {
		return errDBNil
	}

	query := `UPDATE products SET name=$1, price=$2 WHERE id=$3`

	_, err := db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, id string) error {
	if db == nil {
		return errDBNil
	}

	query := `DELETE FROM products WHERE id=$1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func SelectProductInByIDs(db *sql.DB, ids []string) ([]Product, error) {
	if db == nil {
		return nil, errDBNil
	}

	placeholders := []string{}
	arg := []any{}
	for i, id := range ids {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		arg = append(arg, id)
	}

	query := fmt.Sprintf(`SELECT id, name, price FROM products WHERE is_deleted = false AND id IN (%s);`, strings.Join(placeholders, ","))
	rows, err := db.Query(query, arg...)
	if err != nil {
		return nil, err
	}

	products := []Product{}

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil

}
