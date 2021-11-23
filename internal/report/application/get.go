package application

import (
	"deporvillage-feeder-backend/internal/report/domain"
)

type GetApplicationService struct {
	repository domain.ReportRepository
}

type Response struct {
	Uniques    int
	Duplicated int
	Invalids   int
}

func CreateGetApplicationService(repository domain.ReportRepository) GetApplicationService {
	return GetApplicationService{repository}
}

func (as GetApplicationService) Execute() Response {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		return Response{}
	}

	return Response{
		report.GetCounterProduct(),
		report.GetCounterProductDuplicated(),
		report.GetCounterProductInvalid(),
	}
}
