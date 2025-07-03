# Basic Go CLI Application

* Boilerplate code for a Go CLI application 
* Supports environment variable checks and command-line argument parsing.
* Cross platform support (Linux and Windows)
* Light weight 2.5MB binary


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

