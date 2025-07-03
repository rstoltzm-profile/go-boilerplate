package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rstoltzm-profile/go-boilerplate/go-cli-files/internal"
)

func main() {
	// Clean log output, helpful for CLI tools
	log.SetFlags(0)

	// Get Arguments
	variable, inputFile, outFormat, outFile, err := internal.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Check and Set Environment Variables
	err = internal.CheckCustomEnv()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("API_KEY")
	endpoint := os.Getenv("ENDPOINT")

	// Print Variables
	fmt.Printf("Arguments:\nvariable:  %s\ninputFile:  %s\noutFormat:  %s\noutFile:    %s\n",
		variable, inputFile, outFormat, outFile)

	fmt.Printf("\nSecrets:\nAPI_KEY: %s\n", maskSecret(apiKey))
	fmt.Printf("ENDPOINT: %s\n", maskSecret(endpoint))

}

func maskSecret(secret string) string {
	// optional function,
	if len(secret) <= 4 {
		return "****"
	}
	return secret[:2] + "****" + secret[len(secret)-2:]
}
