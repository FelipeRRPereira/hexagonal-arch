package db_test

import (
	"database/sql"
	"testing"

	"github.com/feliperrpereira/go-hexagonal/adapters/db"
	"github.com/feliperrpereira/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db, "123e4567-e89b-12d3-a456-426614174000", "Test Product", 100.0, "disabled")
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		status TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

func createProduct(db *sql.DB, id, name string, price float64, status string) error {
	query := "INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, id, name, price, status)
	return err
}

func TestProductDb_get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(t, err)
	require.Equal(t, "123e4567-e89b-12d3-a456-426614174000", product.GetID())
	require.Equal(t, "Test Product", product.GetName())
	require.Equal(t, 100.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "New Product"
	product.Price = 150.0
	product.Status = "enabled"

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetID(), productResult.GetID())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = "disabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetID(), productResult.GetID())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}
