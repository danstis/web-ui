package main

import (
	"log"

	"github.com/ADO-Asana-Sync/web-ui/internal/version"
)

// Main entry point for the app.
func main() {
	log.Printf("Version %q", version.Version)
}
