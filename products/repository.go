package products

import (
	"log"

	"github.com/Lucashcr/LHShopBackend/dbconn"
)

func GetProductList(page int, itensPerPage int) ([]Product, int) {
	db := dbconn.GetDB()
	defer db.Close()

	countQuery := "SELECT COUNT(id) FROM products"

	var productsCount int
	err := db.QueryRow(countQuery).Scan(&productsCount)
	if err != nil {
		log.Fatal("[DEBUG]: Error querying products count!")
	}

	var products []Product

	query := "SELECT id, name, description, price FROM products LIMIT $1 OFFSET $2"

	limit := itensPerPage
	offset := (page - 1) * itensPerPage

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		log.Fatal("[ERROR]: Error querying products!")
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Fatal("[ERROR]: Error scanning product row!")
		}

		products = append(products, p)
	}

	return products, productsCount
}
