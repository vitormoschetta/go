package repositories

import (
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CategoryRepository struct {
	Db *sql.DB
}

func NewCategoryRepository(db *sql.DB) interfaces.IRepository[models.Category] {
	return &CategoryRepository{Db: db}
}

func (r *CategoryRepository) FindAll() (categories []models.Category, err error) {
	rows, err := r.Db.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Print(err)
			continue
		}
		categories = append(categories, c)
	}
	return
}

func (r *CategoryRepository) FindByID(id string) (category models.Category, err error) {
	row := r.Db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id)
	err = row.Scan(&category.ID, &category.Name)
	if err != nil {
		log.Print(err)
	}
	return
}

func (r *CategoryRepository) Save(p models.Category) error {
	stmt, err := r.Db.Prepare("INSERT INTO categories (id, name) VALUES (?, ?)")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.ID, p.Name)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Category created")
	}
	return nil
}

func (r *CategoryRepository) Update(p models.Category) error {
	stmt, err := r.Db.Prepare("UPDATE categories SET name = ? WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.Name, p.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Category updated")
	}
	return nil
}

func (r *CategoryRepository) Delete(id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM categories WHERE id = ?")
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
		log.Print("Category deleted")
	}
	return nil
}
