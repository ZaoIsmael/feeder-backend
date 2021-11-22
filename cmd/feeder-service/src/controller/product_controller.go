package controller

import (
	"deporvillage-feeder-backend/internal/inventory/application"
	"fmt"
)

type ProductController struct {
	service application.AddProductApplicationService
}

func CreateProductController(service application.AddProductApplicationService) ProductController {
	return ProductController{service}
}

func (c ProductController) Run(i string) {
	err := c.service.Execute(i)

	if err != nil {
		fmt.Println(err)
	}
}
