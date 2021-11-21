package domain

type Product struct {
	Sku SKU
}

func CreateProduct(sku string) (Product, error) {
	SKU, err := CreateSKU(sku)

	if err != nil {
		return Product{}, err
	}

	return Product{SKU}, nil
}
