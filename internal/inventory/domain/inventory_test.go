package domain

import (
	"testing"
)

func TestCreateInventory(t *testing.T) {
	CreateInventory()
}

func TestInventory_AddProduct(t *testing.T) {
	inventory := CreateInventory()
	err := inventory.AddProduct("LPOS-3241")

	if err != nil {
		t.Errorf("should return one product")
	}
}
