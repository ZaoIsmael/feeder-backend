package domain

import "deporvillage-feeder-backend/pkg/domain"

type Report struct {
	domain.AggregateRoot
	Id                      ReportId
	CountProducts           int
	countProductsInvalid    int
	countProductsDuplicated int
}

func CreateReport() Report {
	return Report{Id: ReportId{"1"}, CountProducts: 0, countProductsInvalid: 0, countProductsDuplicated: 0}
}

func (r *Report) IncrementProduct() {
	r.CountProducts++
}

func (r *Report) IncrementProductInvalid() {
	r.countProductsInvalid++
}

func (r *Report) IncrementProductDuplicated() {
	r.countProductsDuplicated++
}
