package cli

import (
	"fmt"

	"github.com/feliperrpereira/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s created successfully with name '%s' and price %.2f", product.GetID(), product.GetName(), product.GetPrice())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		enabledProduct, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s enabled successfully with name '%s' and price %.2f", enabledProduct.GetID(), enabledProduct.GetName(), enabledProduct.GetPrice())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		disabledProduct, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s disabled successfully with name '%s' and price %.2f", disabledProduct.GetID(), disabledProduct.GetName(), disabledProduct.GetPrice())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s has name '%s' and price %.2f", product.GetID(), product.GetName(), product.GetPrice())
	}
	return result, nil
}
