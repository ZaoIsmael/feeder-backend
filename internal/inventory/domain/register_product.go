package domain

import "deporvillage-feeder-backend/internal/cross-cutting/domain"

type RegisterProduct interface {
	Record(sku domain.SKU)
}
