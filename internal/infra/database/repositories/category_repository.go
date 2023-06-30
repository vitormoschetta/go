package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/pkg/utils"
)

type CategoryRepository struct {
	Db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{Db: db}
}

func (r *CategoryRepository) FindAll(ctx context.Context) (categories []category.Category, err error) {
	rows, err := r.Db.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c category.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
			break
		}
		categories = append(categories, c)
	}
	return
}

func (r *CategoryRepository) FindByID(ctx context.Context, id string) (category category.Category, err error) {
	row := r.Db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id)
	errData := row.Scan(&category.ID, &category.Name)
	if errData != nil {
		if errData == sql.ErrNoRows {
			return category, nil
		}
		err = errData
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
	}
	return
}

func (r *CategoryRepository) Save(ctx context.Context, p category.Category) error {
	stmt, err := r.Db.Prepare("INSERT INTO categories (id, name) VALUES (?, ?)")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(p.ID, p.Name)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Category saved"))
	}
	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, p category.Category) error {
	stmt, err := r.Db.Prepare("UPDATE categories SET name = ? WHERE id = ?")
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	res, err := stmt.Exec(p.Name, p.ID)
	if err != nil {
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	if res != nil {
		log.Print(utils.BuildLogger(ctx, "Category updated"))
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM categories WHERE id = ?")
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
		log.Print(utils.BuildLogger(ctx, "Category deleted"))
	}
	return nil
}
