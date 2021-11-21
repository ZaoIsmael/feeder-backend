package application

import (
	"deporvillage-feeder-backend/internal/report/domain"
)

type GetApplicationService struct {
	repository domain.ReportRepository
}

type Response struct {
	total      int
	duplicated int
	invalids   int
}

func CreateGetApplicationService(repository domain.ReportRepository) GetApplicationService {
	return GetApplicationService{repository}
}

func (as GetApplicationService) Execute() (Response, error) {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		return Response{}, err
	}

	return Response{
		report.CountProducts,
		report.CountProductsDuplicated,
		report.CountProductsInvalid,
	}, nil
}
