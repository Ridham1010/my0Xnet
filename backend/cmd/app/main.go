package main

import (
	"log"

	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/db"
	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/discovery"
	httpapi "github.com/bhawani-prajapat2006/0Xnet/backend/internal/http"

	"github.com/google/uuid"
)

func main() {
	deviceID := uuid.New().String()

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	go discovery.Advertise(8080, deviceID)

	server := httpapi.NewServer(dbConn, deviceID)
	log.Println("0Xnet running on port 8080")
	server.Start()
}
