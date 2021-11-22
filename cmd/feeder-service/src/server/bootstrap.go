package server

import (
	"deporvillage-feeder-backend/internal/inventory/application"
	eventHandlerInventory "deporvillage-feeder-backend/internal/inventory/application/event-handlers"
	"deporvillage-feeder-backend/internal/inventory/domain"
	"deporvillage-feeder-backend/internal/inventory/infrastructure"
	applicationReport "deporvillage-feeder-backend/internal/report/application"
	eventHandlerReport "deporvillage-feeder-backend/internal/report/application/event-handlers"
	domainReport "deporvillage-feeder-backend/internal/report/domain"
	infraReport "deporvillage-feeder-backend/internal/report/infrastructure"
	pkgDomain "deporvillage-feeder-backend/pkg/domain"
	pkgInfra "deporvillage-feeder-backend/pkg/infrastructure"
	"strconv"
	"time"
)

type App struct {
	Report  applicationReport.GetApplicationService
	Service application.AddProductApplicationService
}

var filename = "tmp/" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".log"

func Boostrap() (App, error) {
	inventoryRepository := infrastructure.NewInventoryRepository(make(map[string]domain.Inventory))
	reportRepository := infraReport.NewReportRepository(make(map[string]domainReport.Report))
	loggerProduct, err := pkgInfra.NewFileLoggerProduct(filename)

	if err != nil {
		return App{}, err
	}

	handlers := []pkgDomain.EventHandler{
		eventHandlerReport.CreateProductWasAddedApplicationService(reportRepository),
		eventHandlerReport.CreateProductWasInvalidApplicationService(reportRepository),
		eventHandlerReport.CreateProductWasDuplicatedApplicationService(reportRepository),
		eventHandlerInventory.CreateProductWasAddedEventHandler(loggerProduct),
	}

	eventBus := pkgInfra.InMemoryEventBus{Handlers: handlers}

	addProductAS := application.CreateAddProductApplicationService(inventoryRepository, eventBus)
	getReportAS := applicationReport.CreateGetApplicationService(reportRepository)

	return App{
		getReportAS,
		addProductAS,
	}, nil
}
