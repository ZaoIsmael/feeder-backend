package domain

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type Report struct {
	domain.AggregateRoot
	Id                       ReportId
	counterProduct           int
	counterProductInvalid    int
	counterProductDuplicated int
}

type ReportDTO struct {
	Cp int
	Ci int
	Cd int
}

func CreateNewReport() Report {
	return Report{Id: ReportId{"1"}, counterProduct: 0, counterProductInvalid: 0, counterProductDuplicated:0}
}

func CreateReport(r ReportDTO) Report {
	return Report{Id: ReportId{"1"}, counterProduct: r.Cd, counterProductInvalid: r.Ci, counterProductDuplicated: r.Cd}
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

func (r Report) GetCounterProduct() int {
	return r.counterProduct
}

func (r Report) GetCounterProductInvalid() int {
	return r.counterProductInvalid
}

func (r Report) GetCounterProductDuplicated() int {
	return r.counterProductDuplicated
}
