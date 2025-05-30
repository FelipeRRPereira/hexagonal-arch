package application_test

import (
	"testing"

	"github.com/feliperrpereira/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 10.0
	product.Status = application.DISABLED
	product.ID = "12345"

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0.0
	err = product.Enable()
	require.Equal(t, err.Error(), "product price must be greater than 0 to enable")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 10.0
	product.Status = application.ENABLED
	product.ID = "12345"

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 0.0
	err = product.Disable()
	require.Equal(t, err.Error(), "product price must be greater than 0 to disable")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 10.0
	product.Status = application.ENABLED
	product.ID = uuid.NewV4().String()

	valid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, valid)

	product.Name = ""
	valid, err = product.IsValid()
	require.Equal(t, err.Error(), "name is required")
	require.False(t, valid)

	product.Name = "Test Product"
	product.Price = -1.0
	valid, err = product.IsValid()
	require.Equal(t, err.Error(), "product price must be greater than or equal to 0")
	require.False(t, valid)

	product.Price = 10.0
	product.Status = "invalid"
	valid, err = product.IsValid()
	require.Equal(t, err.Error(), "invalid product status")
	require.False(t, valid)
}
