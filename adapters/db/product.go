package db

import (
	"database/sql"

	"github.com/feliperrpereira/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	query := "SELECT id, name, price, status FROM products WHERE id = ?"
	err := p.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	query := "SELECT COUNT(*) FROM products WHERE id = ?"
	err := p.db.QueryRow(query, product.GetID()).Scan(&rows)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return p.create(product)
	}
	return p.update(product)
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}
