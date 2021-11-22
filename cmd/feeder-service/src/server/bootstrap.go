package server

import (
	"deporvillage-feeder-backend/cmd/feeder-service/src/controller"
	pkgDomain "deporvillage-feeder-backend/internal/cross-cutting/domain"
	infrastructure2 "deporvillage-feeder-backend/internal/cross-cutting/infrastructure"
	"deporvillage-feeder-backend/internal/inventory/application"
	eventHandlerInventory "deporvillage-feeder-backend/internal/inventory/application/event-handlers"
	"deporvillage-feeder-backend/internal/inventory/domain"
	"deporvillage-feeder-backend/internal/inventory/infrastructure"
	applicationReport "deporvillage-feeder-backend/internal/report/application"
	eventHandlerReport "deporvillage-feeder-backend/internal/report/application/event-handlers"
	domainReport "deporvillage-feeder-backend/internal/report/domain"
	infraReport "deporvillage-feeder-backend/internal/report/infrastructure"
	"strconv"
	"time"
)

type App struct {
	Report  controller.ReportController
	Product controller.ProductController
}

var filename = "tmp/" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".log"

func Boostrap() (App, error) {
	inventoryRepository := infrastructure.NewInventoryRepository(make(map[string]domain.Inventory))
	reportRepository := infraReport.NewReportRepository(make(map[string]domainReport.Report))
	loggerProduct, err := infrastructure2.NewFileLoggerProduct(filename)

	if err != nil {
		return App{}, err
	}

	handlers := []pkgDomain.EventHandler{
		eventHandlerReport.CreateProductWasAddedApplicationService(reportRepository),
		eventHandlerReport.CreateProductWasInvalidApplicationService(reportRepository),
		eventHandlerReport.CreateProductWasDuplicatedApplicationService(reportRepository),
		eventHandlerInventory.CreateProductWasAddedEventHandler(loggerProduct),
	}

	eventBus := infrastructure2.InMemoryEventBus{Handlers: handlers}

	addProductAS := application.CreateAddProductApplicationService(inventoryRepository, eventBus)
	getReportAS := applicationReport.CreateGetApplicationService(reportRepository)

	return App{
		controller.CreateReportController(getReportAS),
		controller.CreateProductController(addProductAS),
	}, nil
}
