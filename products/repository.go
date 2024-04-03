package products

import (
	"log"

	"github.com/Lucashcr/LHShopBackend/dbconn"
)

func FetchProductList(page int, itensPerPage int) ([]Product, int) {
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

func FetchProductByID(id string) (*Product, error) {
	db := dbconn.GetDB()
	defer db.Close()

	query := "SELECT id, name, description, price FROM products WHERE id = $1"

	var p Product
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price)
	if err != nil {
		log.Println("[ERROR]: Error querying product!")
		return nil, err
	}

	return &p, nil
}

func InsertProductIntoDatabase(p *Product) error {
	db := dbconn.GetDB()
	defer db.Close()

	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("[ERROR]: Error preparing insert statement!")
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.Name, p.Description, p.Price).Scan(&p.ID)
	if err != nil {
		log.Println("[ERROR]: Error inserting product!")
		return err
	}

	return nil
}

func UpdateProductByID(p *Product) error {
	db := dbconn.GetDB()
	defer db.Close()

	query := "UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("[ERROR]: Error preparing update statement!")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.ID)
	if err != nil {
		log.Println("[ERROR]: Error updating product!")
		return err
	}

	return nil
}

func DeleteProductByID(id string) error {
	db := dbconn.GetDB()
	defer db.Close()

	query := "DELETE FROM products WHERE id = $1"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("[ERROR]: Error preparing delete statement!")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("[ERROR]: Error deleting product!")
		return err
	}

	return nil
}
