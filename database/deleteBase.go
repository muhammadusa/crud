package database

import (
	"app/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DeleteProducts(body DeleteProductBody) Product {

	db, err := sql.Open("postgres", config.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("delete 16 %v", err)
	}

	var dropProduct Product

	err = db.QueryRow(SQL_DELETE_QUERY, body.Id).Scan(
		&dropProduct.Id,
		&dropProduct.Name,
		&dropProduct.Price,
		&dropProduct.Category.Id,
	)

	if err != nil {
		log.Fatalf("delete 30 %v", err)
	}

	return dropProduct
}
