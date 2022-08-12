package database

import (
	"app/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func PostProducts(body PostProductBody) Product {

	db, err := sql.Open("postgres", config.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("post 16 %v", err)
	}

	var postProducts Product

	err = db.QueryRow(
		SQL_POST_QUERY,
		body.Name,
		body.Price,
		body.CategoryId,
	).Scan(
		&postProducts.Id,
		&postProducts.Name,
		&postProducts.Price,
		&postProducts.Category.Id,
		// &postProducts.Category.Name,
	)

	log.Println(postProducts)

	if err != nil {
		log.Fatalf("post 35 %v", err)
	}

	return postProducts

}
