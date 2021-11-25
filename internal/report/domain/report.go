package domain

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type Report struct {
	domain.AggregateRoot
	id                       ReportId
	counterProduct           int
	counterProductInvalid    int
	counterProductDuplicated int
}

func CreateReport() Report {
	return Report{
		id: ReportId{"1"},
	}
}

func (r *Report) IncrementProduct() {
	r.counterProduct++
}

func (r *Report) IncrementProductInvalid() {
	r.counterProductInvalid++
}

func (r *Report) IncrementProductDuplicated() {
	r.counterProductDuplicated++
}

func (r Report) Id() ReportId {
	return r.id
}

func (r Report) GetCounterProduct() int {
	return r.counterProduct
}

func (r Report) GetCounterProductInvalid() int {
	return r.counterProductInvalid
}

func (r Report) GetCounterProductDuplicated() int {
	return r.counterProductDuplicated
}
