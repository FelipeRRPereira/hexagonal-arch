package application_test

import (
	"testing"

	"github.com/feliperrpereira/go-hexagonal/application"
	mock_application "github.com/feliperrpereira/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persister := mock_application.NewMockProductPersisterInterface(ctrl)
	persister.EXPECT().Get(gomock.Any()).Return(&product, nil).AnyTimes()

	service := application.ProductService{
		ProductPersister: persister,
	}

	result, err := service.Get("123e4567-e89b-12d3-a456-426614174000")
	require.Nil(t, err)
	require.Equal(t, &product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := mock_application.NewMockProductInterface(ctrl)
	var product application.ProductInterface = productMock
	persister := mock_application.NewMockProductPersisterInterface(ctrl)
	persister.EXPECT().Save(gomock.Any()).Return(&product, nil).AnyTimes()

	service := application.ProductService{
		ProductPersister: persister,
	}

	result, err := service.Create("Test Product", 100.0)
	require.Nil(t, err)
	require.Equal(t, &product, result)
}
func TestProductService_Create_InvalidProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persister := mock_application.NewMockProductPersisterInterface(ctrl)

	service := application.ProductService{
		ProductPersister: persister,
	}

	// Attempt to create a product with an invalid name
	result, err := service.Create("", 100.0)
	require.NotNil(t, err)
	require.Nil(t, result)
	require.Equal(t, "name is required", err.Error())
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := mock_application.NewMockProductInterface(ctrl)
	var product application.ProductInterface = productMock
	persister := mock_application.NewMockProductPersisterInterface(ctrl)
	persister.EXPECT().Save(gomock.Any()).Return(&product, nil).AnyTimes()

	service := application.ProductService{
		ProductPersister: persister,
	}

	// Enable the product
	productMock.EXPECT().Enable().Return(nil).AnyTimes()
	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, &product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := mock_application.NewMockProductInterface(ctrl)
	var product application.ProductInterface = productMock
	persister := mock_application.NewMockProductPersisterInterface(ctrl)
	persister.EXPECT().Save(gomock.Any()).Return(&product, nil).AnyTimes()

	service := application.ProductService{
		ProductPersister: persister,
	}

	// Disable the product
	productMock.EXPECT().Disable().Return(nil).AnyTimes()
	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, &product, result)
}
