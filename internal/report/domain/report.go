package domain

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type Report struct {
	domain.AggregateRoot
	Id                      ReportId
	CountProducts           int
	CountProductsInvalid    int
	CountProductsDuplicated int
}

func CreateReport() Report {
	return Report{Id: ReportId{"1"}, CountProducts: 0, CountProductsInvalid: 0, CountProductsDuplicated: 0}
}

func (r *Report) IncrementProduct() {
	r.CountProducts++
}

func (r *Report) IncrementProductInvalid() {
	r.CountProductsInvalid++
}

func (r *Report) IncrementProductDuplicated() {
	r.CountProductsDuplicated++
}
