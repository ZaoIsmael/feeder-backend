package domain

type ReportDTO struct {
	Id                       string
	CounterProduct           int
	CounterProductInvalid    int
	CounterProductDuplicated int
}

func (r ReportDTO) ToDomain() Report {
	return Report{
		id:                       ReportId{r.Id},
		counterProduct:           r.CounterProduct,
		counterProductInvalid:    r.CounterProductInvalid,
		counterProductDuplicated: r.CounterProductDuplicated,
	}
}
