package internal

import (
	"flag"
	"fmt"
	"os"
)

// Note: Update constants as needed
const (
	flagVariable     = "variable"
	flagVariableDesc = "variable name"

	flagInputFile     = "inputFile"
	flagInputFileDesc = "filename for inputFile"

	flagOutFormat     = "outFormat"
	flagOutFormatDesc = "output type: json"

	flagOutFile     = "outFile"
	flagOutFileDesc = "output file name"

	errorApiKey   = "API_KEY env variable is not set"
	errorEndpoint = "ENDPOINT env variable is not set"
)

func PrintUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]\n\n", os.Args[0])
	// Note: Update required variables for flags as needed
	fmt.Fprintln(flag.CommandLine.Output(), "Required flags:")
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s   string   %s\n", flagVariable, flagVariableDesc)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  string   %s\n", flagInputFile, flagInputFileDesc)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  string   %s\n", flagOutFormat, flagOutFormatDesc)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s    string   %s\n", flagOutFile, flagOutFileDesc)

	// Note: Update required env variables as needed
	fmt.Fprintln(flag.CommandLine.Output(), "\nRequired environment variables:")
	fmt.Fprintln(flag.CommandLine.Output(), "  API_KEY     your API key")
	fmt.Fprintln(flag.CommandLine.Output(), "  ENDPOINT    your endpoint URL")
	fmt.Fprintln(flag.CommandLine.Output(), "\nExample:")

	// Note: Update example message as needed
	fmt.Fprintf(flag.CommandLine.Output(), "  %s -%s foo -%s in.json -%s json -%s out.json\n",
		os.Args[0], flagVariable, flagInputFile, flagOutFormat, flagOutFile)
}

func ParseFlags() (variable string, inputFile string, outType string, outputFile string, err error) {
	// Set flag Usage
	flag.Usage = PrintUsage

	// Show help if no arguments are provided
	if len(os.Args) == 1 {
		flag.Usage()
		return "", "", "", "", fmt.Errorf("no arguments provided")
	}

	v := flag.String(flagVariable, "", flagVariableDesc)
	i := flag.String(flagInputFile, "", flagInputFileDesc)
	o := flag.String(flagOutFormat, "", flagOutFormatDesc)
	of := flag.String(flagOutFile, "", flagOutFileDesc)

	flag.Parse()

	// Print usage and exit if any required flag is missing
	missing := false
	if *v == "" || *i == "" || *o == "" || *of == "" {
		missing = true
	}
	if missing {
		flag.Usage()
		return *v, *i, *o, *of, fmt.Errorf("missing required flag(s)")
	}

	return *v, *i, *o, *of, nil
}

func CheckCustomEnv() error {
	if os.Getenv("API_KEY") == "" {
		return fmt.Errorf(errorApiKey)
	}
	if os.Getenv("ENDPOINT") == "" {
		return fmt.Errorf(errorEndpoint)
	}
	return nil
}
