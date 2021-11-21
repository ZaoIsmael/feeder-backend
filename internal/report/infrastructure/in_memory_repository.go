package infrastructure

import (
	"deporvillage-feeder-backend/internal/report/domain"
	"errors"
)

type ReportRepository struct {
	R map[string]domain.Report
}

func NewReportRepository(r map[string]domain.Report) ReportRepository {
	return ReportRepository{
		R: r,
	}
}

func (r ReportRepository) Find(id domain.ReportId) (domain.Report, error) {
	i, ok := r.R[id.Value]

	if ok {
		return i, nil
	}

	return domain.Report{}, errors.New("the report not exist")
}

func (r *ReportRepository) Save(re domain.Report) {
	r.R[re.Id.Value] = re
}
