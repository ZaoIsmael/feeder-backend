package infrastructure

import (
	"deporvillage-feeder-backend/internal/report/domain"
	"errors"
)

type ReportRepository struct {
	r map[string]domain.Report
}

func NewReportRepository(r map[string]domain.Report) ReportRepository {
	return ReportRepository{
		r: r,
	}
}

func (r ReportRepository) Find(id domain.ReportId) (domain.Report, error) {
	i, ok := r.r[id.Value]

	if ok {
		return i, nil
	}

	return domain.Report{}, errors.New("the report not exist")
}

func (r ReportRepository) Save(re domain.Report) {
	i, ok := r.r[re.Id.Value]

	if ok {
		return
	}

	r.r[re.Id.Value] = i
}
