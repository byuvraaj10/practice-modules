package repository

import (
	"database/sql"
	"ecommerce-inventory/model"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Add a product
func (repo *ProductRepository) AddProduct(product *model.Product) error {
	_, err := repo.db.Exec(`INSERT INTO products (name, description, price, stock, category_id) 
		VALUES (?, ?, ?, ?, ?)`, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	return err
}

// Get a product by ID
func (repo *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	row := repo.db.QueryRow(`SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`, id)
	product := &model.Product{}
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// Update a product
func (repo *ProductRepository) UpdateProduct(product *model.Product) error {
	_, err := repo.db.Exec(`UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? 
		WHERE id = ?`, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ID)
	return err
}

// Delete a product
func (repo *ProductRepository) DeleteProduct(id int) error {
	_, err := repo.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}

// Get all products with pagination
func (repo *ProductRepository) GetAllProducts(page, limit int) ([]model.Product, error) {
	rows, err := repo.db.Query(`SELECT id, name, description, price, stock, category_id FROM products LIMIT ? OFFSET ?`,
		limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
