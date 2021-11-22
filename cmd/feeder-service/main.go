package main

import (
	"deporvillage-feeder-backend/cmd/feeder-service/src/server"
	"log"
)

func main() {
	app, err := server.Boostrap()

	if err != nil {
		log.Fatal(err)
		return
	}

	srv, err := server.CreateServer(app.Product)

	if err != nil {
		log.Fatal(err)
		return
	}

	srv.Run()
	srv.Shutdown()

	app.Report.Run("")
}
