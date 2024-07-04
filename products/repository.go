package products

import (
	"log"

	"github.com/Lucashcr/LHShopBackend/dbconn"
)

func fetchProductList(page int, itensPerPage int) ([]product, int, error) {
	db := dbconn.GetDB()

	var products []product
	var productsCount int

	countQuery := "SELECT COUNT(id) FROM products"

	err := db.QueryRow(countQuery).Scan(&productsCount)
	if err != nil {
		log.Println("[DEBUG]: Error querying products count!")
		return []product{}, 0, err
	}

	query := "SELECT id, name, description, price FROM products LIMIT $1 OFFSET $2"

	limit := itensPerPage
	offset := (page - 1) * itensPerPage

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		log.Println("[ERROR]: Error querying products!")
		return []product{}, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Println("[ERROR]: Error scanning product row!")
			return []product{}, productsCount, err
		}

		products = append(products, p)
	}

	return products, productsCount, nil
}

func fetchProductByID(id string) (*product, error) {
	db := dbconn.GetDB()

	query := "SELECT id, name, description, price FROM products WHERE id = $1"

	var p product
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price)
	if err != nil {
		log.Println("[ERROR]: Error querying product!")
		return nil, err
	}

	return &p, nil
}

func insertProductIntoDatabase(p *product) error {
	db := dbconn.GetDB()

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

func updateProductByID(p *product) (int64, error) {
	db := dbconn.GetDB()

	query := "UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("[ERROR]: Error preparing update statement!")
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Description, p.Price, p.ID)
	if err != nil {
		log.Println("[ERROR]: Error updating product!")
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("[ERROR]: Error getting rows affected!")
		return 0, err
	}

	return rowsAffected, nil
}

func deleteProductByID(id string) (int64, error) {
	db := dbconn.GetDB()

	query := "DELETE FROM products WHERE id = $1"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("[ERROR]: Error preparing delete statement!")
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println("[ERROR]: Error deleting product!")
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("[ERROR]: Error getting rows affected!")
		return 0, err
	}

	return rowsAffected, nil
}
