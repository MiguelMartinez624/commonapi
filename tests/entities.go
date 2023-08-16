package tests

import "github.com/miguelmartinez624/commonapi/common"

type ProductSearchParams struct {
	common.Filter
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Price struct {
	// TODO if not changes set use the field Name
	Amount   float64 `json:"amount" changes:"Amount"`
	Currency string  `json:"currency" changes:"Amount"`
}

type Tag struct {
	Key string `json:"key" form:"Key"`
	// Why any?
	Value any `json:"value" form:"value"`
}
type ProductInput struct {
	Name        string `json:"name" form:"Name" changes:"Name"`
	Description string `json:"description" form:"Description" changes:"Description"`
	SKU         string `json:"sku" form:"SKU" changes:"SKU"`
	Price       Price  `json:"price" changes:"Price"`
	Tags        []Tag  `json:"tags" form:"Tags" changes:"Tags"`
}

type ProductEntity struct {
	ID             string   `json:"id"`
	Name           string   `json:"name" listing:"Name"`
	Description    string   `json:"description" listing:"Description"`
	SKU            string   `json:"sku" form:"SKU" changes:"SKU"`
	Price          Price    `json:"price"`
	Tags           []Tag    `json:"tags"`
	EventsOnEntity []string `json:"eventsOnEntity"`
}
