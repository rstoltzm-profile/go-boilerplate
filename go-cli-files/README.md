# Basic Go CLI Application

* Boilerplate code for a Go CLI application 
* Supports environment variable checks and command-line argument parsing.
* Cross platform support (Linux and Windows)
* Light weight binary ~2.5MB

## Use Cases
* Can be extended to read the filename passed in, process it, and output results in various formats.
* Use inputs and env variables to call an API.
* Change inputs and env variables to support any CLI use case.

## Usage
### Linux
```bash
export API_KEY="blah"
export ENDPOINT="blah.com"
bin/cli-template --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
```

### Powershell
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

