package domain

import (
	"errors"
	"regexp"
)

type SKU struct {
	Value string
}

var SkuInvalidError = errors.New("the product SKU is invalid")

func CreateSKU(value string) (SKU, error) {
	err := validate(value)

	if err != nil {
		return SKU{}, err
	}

	return SKU{value}, nil
}

func validate(v string) error {
	if regexp.MustCompile(`^[A-Z]{4}-[0-9]{4}$`).MatchString(v) {
		return nil
	}

	return SkuInvalidError
}
