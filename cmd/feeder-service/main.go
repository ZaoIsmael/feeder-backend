package main

import (
	server2 "deporvillage-feeder-backend/cmd/feeder-service/src/server"
	"fmt"
	"log"
)

func main() {
	app, err := server2.Boostrap()

	if err != nil {
		log.Fatal(err)
		return
	}

	srv, err := server2.CreateServer(app.Service)

	if err != nil {
		log.Fatal(err)
		return
	}

	srv.Run()
	srv.Shutdown()

	re, _ := app.Report.Execute()

	fmt.Println(re)
}
