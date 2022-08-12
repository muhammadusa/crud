package database

import (
	"app/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetProducts() []Product {

	db, err := sql.Open("postgres", config.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("get 27 = %v", err)
	}

	rows, err := db.Query(SQL_GET_QUERY)

	if err != nil {
		log.Fatalf("get 47 = %v", err)
	}

	var products []Product

	for rows.Next() {

		var product Product

		rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Category.Id,
			&product.Category.Name,
		)

		products = append(products, product)
	}

	return products
}
