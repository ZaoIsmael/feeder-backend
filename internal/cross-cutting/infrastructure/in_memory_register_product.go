package infrastructure

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type InMemoryRegisterProduct struct {
	d map[string]struct{}
}

func CreateInMemoryRegisterProduct(d map[string]struct{}) *InMemoryRegisterProduct {
	return &InMemoryRegisterProduct{
		d,
	}
}

func (r *InMemoryRegisterProduct) Record(sku domain.SKU) {
	_, ok := r.d[sku.Value()]

	if ok {
		return
	}

	r.d[sku.Value()] = struct{}{}
}

func (r InMemoryRegisterProduct) Exist(sku domain.SKU) bool {
	_, ok := r.d[sku.Value()]

	return ok
}
