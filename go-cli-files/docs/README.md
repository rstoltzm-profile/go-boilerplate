# Documentation for Go CLI Application


## Build and Test outputs
### Make Test
```text
go test ./internal -v
=== RUN   TestEnvVariables
=== RUN   TestEnvVariables/missing_ENDPOINT
=== RUN   TestEnvVariables/missing_API_KEY
=== RUN   TestEnvVariables/all_present
--- PASS: TestEnvVariables (0.00s)
    --- PASS: TestEnvVariables/missing_ENDPOINT (0.00s)
    --- PASS: TestEnvVariables/missing_API_KEY (0.00s)
    --- PASS: TestEnvVariables/all_present (0.00s)
=== RUN   TestParseFlags
=== RUN   TestParseFlags/test_all_Inputs
=== RUN   TestParseFlags/test_missing_input
Usage: cmd [OPTIONS]

Required flags:
  -variable   string   variable name
  -inputFile  string   filename for inputFile
  -outFormat  string   output type: json
  -outFile    string   output file name

Required environment variables:
  API_KEY     your API key
  ENDPOINT    your endpoint URL

Example:
  cmd -variable foo -inputFile in.json -outFormat json -outFile out.json
--- PASS: TestParseFlags (0.00s)
    --- PASS: TestParseFlags/test_all_Inputs (0.00s)
    --- PASS: TestParseFlags/test_missing_input (0.00s)
PASS
ok      github.com/rstoltzm-profile/go-boilerplate/go-cli-files/internal
```

### Make build
```text
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/cli-template main.go
GOOS=windows GOARCH=amd64 go build -o bin/cli-template.exe main.go
```

### Make integ-test
```text
## 1. test no arguments ##

Usage: bin/cli-template [OPTIONS]

Required flags:
  -variable   string   variable name
  -inputFile  string   filename for inputFile
  -outFormat  string   output type: json
  -outFile    string   output file name

Required environment variables:
  API_KEY     your API key
  ENDPOINT    your endpoint URL

Example:
  bin/cli-template -variable foo -inputFile in.json -outFormat json -outFile out.json
no arguments provided

## 2. test no env api_key ##

API_KEY env variable is not set

## 3. test no env endpoint ##

ENDPOINT env variable is not set

## 4. test cli ##

Arguments:
variable:  test
inputFile:  file.json
outFormat:  json
outFile:    outFile.json

Secrets:
API_KEY: ****
ENDPOINT: bl****om
```

