package infrastructure

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type FileLoggerProduct struct {
	f *os.File
}

func (fl FileLoggerProduct) listenSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func(f FileLoggerProduct) {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			f.close()
		case syscall.SIGTERM:
			f.close()
		}
	}(fl)
}

func NewFileLoggerProduct(fileName string) (*FileLoggerProduct, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
		return &FileLoggerProduct{}, err
	}

	fl := &FileLoggerProduct{
		f,
	}

	fl.listenSignals()

	return fl, nil
}

func (fl FileLoggerProduct) Record(sku domain.SKU) {
	defer fl.flush()
	_, err := fl.f.WriteString(sku.Value + "\n")

	if err != nil {
		log.Println(err)
	}
}

func (fl FileLoggerProduct) close() {
	err := fl.f.Close()
	if err != nil {
		log.Println(err)
	}
}

func (fl FileLoggerProduct) flush() {
	err := fl.f.Sync()
	if err != nil {
		log.Println(err)
	}
}
