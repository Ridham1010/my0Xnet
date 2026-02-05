package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/db"
	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/discovery"
	httpapi "github.com/bhawani-prajapat2006/0Xnet/backend/internal/http"

	"github.com/google/uuid"
)

func main() {
	deviceID := uuid.New().String()
	
	// Get port from environment variable, default to 8080
	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	log.Printf("Starting 0Xnet on port %d with device ID: %s", port, deviceID)

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start mDNS advertisement with context
	go discovery.Advertise(ctx, port, deviceID)

	// Initialize session discovery
	sessionDiscovery := discovery.NewSessionDiscovery(deviceID)
	sessionDiscovery.StartDiscovery()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-sigChan
		log.Println("\nReceived shutdown signal, cleaning up...")
		cancel()
		os.Exit(0)
	}()

	server := httpapi.NewServer(dbConn, deviceID, sessionDiscovery, port)
	log.Printf("0Xnet running on port %d", port)
	log.Println("Press Ctrl+C to stop")
	server.Start()
}
