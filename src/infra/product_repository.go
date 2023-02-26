package infra

import (
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/src/domain/interfaces"
	"github.com/vitormoschetta/go/src/domain/models"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) interfaces.ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	rows, err := r.Db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (models.Product, error) {
	var p models.Product
	row := r.Db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		log.Fatal(err)
	}
	return p, nil
}

func (r *ProductRepository) Create(p models.Product) (models.Product, error) {
	stmt, err := r.Db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		log.Println("Product created")
	}

	return p, nil
}

func (r *ProductRepository) Update(p models.Product) (models.Product, error) {
	stmt, err := r.Db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		log.Println("Product updated")
	}

	return p, nil
}

func (r *ProductRepository) Delete(id int) error {
	stmt, err := r.Db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil {
		log.Println("Product deleted")
	}

	return nil
}
