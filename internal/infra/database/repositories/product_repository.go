package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/shared/utils"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) FindAll(ctx context.Context) (products []product.Product, err error) {
	query := "SELECT p.id, p.name, p.price, c.id, c.name "
	query += "FROM products p "
	query += "INNER JOIN categories c ON p.category_id = c.id"
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p product.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category.ID, &p.Category.Name)
		if err != nil {
			log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
			continue
		}
		products = append(products, p)
	}
	return
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (product product.Product, err error) {
	row := r.Db.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = ?", id)
	errData := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category.ID)
	if errData != nil {
		if errData == sql.ErrNoRows {
			return
		}
		err = errData
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	return
}

func (r *ProductRepository) FindByCategory(ctx context.Context, categoryID string) (products []product.Product, err error) {
	rows, errData := r.Db.Query("SELECT id, name, price, category_id FROM products WHERE category_id = ?", categoryID)
	if errData != nil {
		if errData == sql.ErrNoRows {
			return
		}
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p product.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
			continue
		}
		products = append(products, p)
	}
	return
}

func (r *ProductRepository) Save(ctx context.Context, p product.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products (id, name, price, category_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(p.ID, p.Name, p.Price, p.Category.ID)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Product saved"))
	}
	return nil
}

func (r *ProductRepository) Update(ctx context.Context, p product.Product) error {
	stmt, err := r.Db.Prepare("UPDATE products SET name = ?, price = ?, category_id = ? WHERE id = ?")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(p.Name, p.Price, p.Category.ID, p.ID)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Product updated"))
	}
	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Product deleted"))
	}
	return nil
}

func (r *ProductRepository) ApplyPromotionOnProductsByCategory(ctx context.Context, categoryId string, percentage float64) error {
	stmt, err := r.Db.Prepare("UPDATE products SET price = price - (price * ?) WHERE category_id = ?")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(percentage, categoryId)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Promotion applied"))
	}
	return nil
}
