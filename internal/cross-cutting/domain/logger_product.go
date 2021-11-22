package domain

type LoggerProduct interface {
	Record(sku SKU)
}
