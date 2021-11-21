package domain

import (
	"testing"
)

func TestCreateInventory(t *testing.T) {
	CreateInventory()
}

func TestInventory_AddProduct(t *testing.T) {
	inventory := CreateInventory()
	inventory.AddProduct("LPOS-3241")

	_, ok := inventory.products["LPOS-3241"]

	if !ok {
		t.Errorf("should return one product")
	}
}
