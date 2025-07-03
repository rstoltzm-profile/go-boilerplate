# go-cli-files

* Boilerplate code for a Go CLI application with environment variable checks and command-line argument parsing.

## Usage

```bash
export API_KEY="blah"
export ENDPOINT="blah.com"
bin/cli-template --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
```

```powershell
$env:API_KEY = "blah"
$env:ENDPOINT = "blah.com"
& bin/cli-template.exe --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
```

## Build and Test
```
make test
make build
make integ-test
```

