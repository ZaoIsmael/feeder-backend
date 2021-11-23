package infrastructure

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type FileRegisterProduct struct {
	f *os.File
}

func (fl FileRegisterProduct) listenSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func(f FileRegisterProduct) {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			f.close()
		case syscall.SIGTERM:
			f.close()
		}
	}(fl)
}

func NewFileRegisterProduct(fileName string) (*FileRegisterProduct, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return &FileRegisterProduct{}, err
	}

	fl := &FileRegisterProduct{
		f,
	}

	fl.listenSignals()

	return fl, nil
}

func (fl FileRegisterProduct) Record(sku domain.SKU) {
	defer fl.flush()
	_, err := fl.f.WriteString(sku.Value() + "\n")

	if err != nil {
		log.Println(err)
	}
}

func (fl FileRegisterProduct) close() {
	err := fl.f.Close()
	if err != nil {
		log.Println(err)
	}
}

func (fl FileRegisterProduct) flush() {
	err := fl.f.Sync()
	if err != nil {
		log.Println(err)
	}
}
