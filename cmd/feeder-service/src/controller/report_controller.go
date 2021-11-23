package controller

import (
	"deporvillage-feeder-backend/internal/report/application"
	"fmt"
)

type ReportController struct {
	service application.GetApplicationService
}

func CreateReportController(service application.GetApplicationService) ReportController {
	return ReportController{service}
}

func (c ReportController) Run(_ string) {
	report, err := c.service.Execute()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received %d unique product skus, %d duplicates and %d discarded values.\n",
		report.Uniques,
		report.Duplicated,
		report.Invalids,
	)
}
