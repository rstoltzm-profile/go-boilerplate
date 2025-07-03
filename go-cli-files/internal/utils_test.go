package internal

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestEnvVariables(t *testing.T) {
	os.Unsetenv("API_KEY")
	os.Unsetenv("ENDPOINT")

	t.Run("missing ENDPOINT", func(t *testing.T) {
		os.Setenv("API_KEY", "dummy")
		os.Unsetenv("ENDPOINT")
		got := CheckCustomEnv()
		want := fmt.Errorf(errorEndpoint)
		if got == nil || got.Error() != want.Error() {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("missing API_KEY", func(t *testing.T) {
		os.Unsetenv("API_KEY")
		os.Setenv("ENDPOINT", "http://localhost:9200")
		got := CheckCustomEnv()
		want := fmt.Errorf(errorApiKey)
		if got == nil || got.Error() != want.Error() {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("all present", func(t *testing.T) {
		os.Setenv("API_KEY", "dummy")
		os.Setenv("ENDPOINT", "http://localhost:9200")
		got := CheckCustomEnv()
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}

func TestParseFlags(t *testing.T) {
	t.Run("test all Inputs", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"cmd",
			"--variable=test",
			"--inputFile=file.json",
			"--outFormat=json",
			"--outFile=outFile.json",
		}
		variable, inputFile, outFormat, outFile, err := ParseFlags()
		if err != nil {
			t.Errorf("expected nil, got error")
		}
		if variable != "test" {
			t.Errorf("expected repo 'test', got '%s'", variable)
		}
		if inputFile != "file.json" {
			t.Errorf("expected tag 'file.json', got '%s'", inputFile)
		}
		if outFormat != "json" {
			t.Errorf("expected asset-name 'json', got '%s'", outFormat)
		}
		if outFile != "outFile.json" {
			t.Errorf("expected output 'outFile.json', got '%s'", outFile)
		}
	})

	t.Run("test missing input", func(t *testing.T) {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"cmd",
			"--variable=test",
			"--inputFile=file.json",
			"--outFormat=json",
		}
		_, _, _, outFile, err := ParseFlags()
		if err == nil {
			t.Errorf("expected error for missing required flag, got nil")
		}
		if outFile != "" {
			t.Errorf("expected output '', got '%s'", outFile)
		}
	})
}
