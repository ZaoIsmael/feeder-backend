package domain

import (
	"testing"
)

func TestCreateInventory(t *testing.T) {
	_, _ = CreateInventory([]string{})
}

func TestInventory_AddProduct(t *testing.T) {
	inventory, _ := CreateInventory([]string{})
	err := inventory.AddProduct("LPOS-3241")

	if err != nil {
		t.Errorf("should return one product")
	}
}
