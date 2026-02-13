package main

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	apiURL := os.Getenv("SATOGRAM_API_URL")
	if apiURL == "" {
		apiURL = "https://api.satogram.xyz"
	}

	client := NewSatogramClient(apiURL)

	srv := server.NewMCPServer(
		"satogram-mcp",
		"1.0.0",
	)

	registerTools(srv, client)

	if err := server.ServeStdio(srv); err != nil {
		fmt.Fprintf(os.Stderr, "satogram-mcp: %v\n", err)
		os.Exit(1)
	}
}
