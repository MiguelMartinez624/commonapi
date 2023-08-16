package tests

import (
	"github.com/miguelmartinez624/commonapi"
	"github.com/miguelmartinez624/commonapi/errors"
	"reflect"
	"testing"
)

func successSaveFunc[T any](entity T) (*T, error) {
	return &entity, nil
}

func appendEvent(product *ProductEntity) *errors.CommandError {
	product.EventsOnEntity = append(product.EventsOnEntity, "First Product")
	return nil
}

func appendSecondEvent(product *ProductEntity) *errors.CommandError {
	product.EventsOnEntity = append(product.EventsOnEntity, "Second Product")
	return nil
}

func TestCreateCommand(t *testing.T) {
	commandDescription := commonapi.CreateDescription[ProductInput, ProductEntity]{}

	command := commandDescription.
		WriteTo(successSaveFunc[ProductEntity]).
		DoBeforeWrite(
			appendEvent,
			appendSecondEvent).
		Done()

	result, err := command.Execute(ProductInput{
		Name:        "Product One",
		Description: "Description del producto",
		SKU:         "Testing",
		Price: Price{
			Amount:   23.5,
			Currency: "USD",
		},
		Tags: nil,
	})

	if err != nil {
		t.Errorf("should not have returned a error")
	}
	expectedResult := ProductEntity{
		Name:        "Product One",
		Description: "Description del producto",
		SKU:         "Testing",
		Price: Price{
			Amount:   23.5,
			Currency: "USD",
		},
		Tags:           nil,
		EventsOnEntity: []string{"First Product", "Second Product"},
	}

	if !reflect.DeepEqual(*result, expectedResult) {
		t.Errorf("%+v is not equal to %+v", result, expectedResult)
	}

}
