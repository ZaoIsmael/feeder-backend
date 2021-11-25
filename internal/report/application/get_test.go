package application

import (
	"deporvillage-feeder-backend/internal/report/domain"
	"deporvillage-feeder-backend/internal/report/infrastructure"
	"testing"
)

func setupTest(d map[string]domain.Report) GetApplicationService {
	repository := infrastructure.NewReportRepository(d)
	return CreateGetApplicationService(repository)
}

func TestGetWhenReportDontExist(t *testing.T) {
	// arrange
	d := map[string]domain.Report{}
	as := setupTest(d)

	r := as.Execute()

	if r.Duplicated > 0 && r.Uniques > 0 && r.Invalids > 0 {
		t.Errorf("error when running application servirce")
	}
}

func TestGetReport(t *testing.T) {
	// arrange
	d := map[string]domain.Report{}
	re := domain.ReportDTO{
		Id:                       "1",
		CounterProduct:           1,
		CounterProductInvalid:    1,
		CounterProductDuplicated: 1,
	}.ToDomain()
	d[re.Id().Value] = re

	as := setupTest(d)

	r := as.Execute()

	if r.Duplicated != 1 && r.Uniques != 1 && r.Invalids != 1 {
		t.Errorf("error when running application servirce")
	}
}
