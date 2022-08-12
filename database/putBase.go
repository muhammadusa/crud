package database

import (
	"app/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func PutProducts(body PutProductBody) Product {

	db, err := sql.Open("postgres", config.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("put 16 %v", err)
	}

	var putProducts Product

	err = db.QueryRow(
		SQL_PUT_QUERY,
		body.Id,
		body.Name,
		body.Price,
		body.CategoryId,
	).Scan(
		&putProducts.Id,
		&putProducts.Name,
		&putProducts.Price,
		&putProducts.Category.Id,
	)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return putProducts
}
