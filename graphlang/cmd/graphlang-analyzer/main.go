package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/riyarana25/code-indexer/graphlang"
)

// createNeo4jDriver creates a Neo4j driver with retry logic.
func createNeo4jDriver() (neo4j.DriverWithContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	retryDelay := time.Second * 5

	neo4jUri := os.Getenv("NEO4J_URI")
	neo4jUser := os.Getenv("NEO4J_USER")
	neo4jPassword := os.Getenv("NEO4J_PASSWORD")

	if neo4jUri == "" || neo4jUser == "" || neo4jPassword == "" {
		return nil, fmt.Errorf("missing required environment variables: NEO4J_URI, NEO4J_USER, NEO4J_PASSWORD")
	}

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to Neo4j after retries: %v", ctx.Err())
		case <-time.After(retryDelay):
			driver, err := neo4j.NewDriverWithContext(
				neo4jUri,
				neo4j.BasicAuth(neo4jUser, neo4jPassword, ""),
			)
			if err == nil {
				if err = driver.VerifyConnectivity(ctx); err == nil {
					fmt.Println("Connected to Neo4j")
					return driver, nil
				}
				driver.Close(ctx)
			}
			fmt.Printf("Retrying connection to Neo4j after %v...\n", retryDelay)
		}
	}
}

func main() {
	driver, err := createNeo4jDriver()
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close(context.Background())

	dirPath := os.Getenv("ANALYZE_DIR")
	if dirPath == "" {
		dirPath = "/app"
	}

	parser := graphlang.NewTreeSitterParser(driver)
	if err := parser.AnalyzeDirectory(dirPath); err != nil {
		log.Fatalf("Analysis failed: %v", err)
	}
}
