package repositories

import (
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/internal/domain/models"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) FindAll() (products []models.Product, err error) {
	rows, err := r.Db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Print(err)
			continue
		}
		products = append(products, p)
	}
	return
}

func (r *ProductRepository) FindByID(id string) (product models.Product, err error) {
	row := r.Db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	err = row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Print(err)
	}
	return
}

func (r *ProductRepository) Save(p models.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product created")
	}
	return nil
}

func (r *ProductRepository) Update(p models.Product) error {
	stmt, err := r.Db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product updated")
	}
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product deleted")
	}
	return nil
}
