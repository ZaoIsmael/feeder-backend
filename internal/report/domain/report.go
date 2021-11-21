package domain

import "deporvillage-feeder-backend/pkg/domain"

type Report struct {
	domain.AggregateRoot
	id                      int
	countProducts           int
	countProductsInvalid    int
	countProductsDuplicated int
}

func CreateReport() Report {
	return Report{id: 1, countProducts: 0, countProductsInvalid: 0, countProductsDuplicated: 0}
}

func (r Report) IncrementProduct() {
	r.countProducts++
}

func (r Report) IncrementProductInvalid() {
	r.countProductsInvalid++
}

func (r Report) IncrementProductDuplicated() {
	r.countProductsDuplicated++
}
