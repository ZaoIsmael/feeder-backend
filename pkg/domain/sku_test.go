package domain

import (
	"testing"
)

var skusValid = []string{
	"KASL-3423", "LPOS-3241",
}

var skusInvalid = []string{
	"LPOS-32411", "123-@", "-fa21-f", "1234-ABCD", "1AB4-1BC4", "A1B4-BBCC",
}

func TestNewProductSKU(t *testing.T) {
	for _, sku := range skusValid {
		t.Run(sku, func(t *testing.T) {
			_, err := CreateSKU(sku)

			if err != nil {
				t.Errorf("Test error")
			}
		})
	}

	for _, sku := range skusInvalid {
		t.Run(sku, func(t *testing.T) {
			_, err := CreateSKU(sku)

			if err == nil {
				t.Errorf("Test error")
			}
		})
	}
}
