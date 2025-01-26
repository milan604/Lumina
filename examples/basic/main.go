package main

import (
	"context"
	"fmt"

	"github.com/milan604/lumina"
)

func main() {
	// Create a new logger
	logger, err := lumina.NewLogger()
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	// Log a message
	logger.Info("Hello, world!")

	// Log a message with fields
	logger.InfoF("Hello, world! %s", "with fields")

	// Log a message with fields and context
	ctx := context.Background()
	ctx = context.WithValue(ctx, lumina.RequestIDKey, "test-request-id")
	logger.InfoFCtx(ctx, "Hello, world! %s", "with fields")
}

// Output:
// {"level":"info","timestamp":"2025-01-26T15:44:21.310+0545","caller":"lumina/logger.go:100","message":"Hello, world!"}
// {"level":"info","timestamp":"2025-01-26T15:44:21.317+0545","caller":"lumina/logger.go:116","message":"Hello, world! with fields"}
// {"level":"info","timestamp":"2025-01-26T15:44:21.317+0545","caller":"lumina/logger.go:132","message":"Hello, world! with fields","requestID":"test-request-id"}
