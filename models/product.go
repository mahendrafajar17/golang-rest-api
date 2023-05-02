package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Qty         int    `json:"qty"`
	LastUpdated string `json:"last_updated"`
}

func GetProducts(db *sql.DB) []Product {
	results, err := db.Query("SELECT * FROM product")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Product{}
	for results.Next() {
		var prod Product
		// for each row, scan into the Product struct
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		products = append(products, prod)
	}

	return products

}

func GetProduct(db *sql.DB, code string) *Product {
	prod := &Product{}
	results, err := db.Query("SELECT * FROM product where code=?", code)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return prod
}

func AddProduct(db *sql.DB, product Product) {
	insert, err := db.Query(
		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
		product.Code, product.Name, product.Qty)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
