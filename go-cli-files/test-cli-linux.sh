#! /usr/bin/bash
echo -e "\n## 1. test no arguments ##\n"
bin/cli-template

echo -e "\n## 2. test no env api_key ##\n"
bin/cli-template --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"

echo -e "\n## 3. test no env endpoint ##\n"
export API_KEY="blah"
bin/cli-template --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
unset API_KEY

echo -e "\n## 4. test cli ##\n"
export API_KEY="blah"
export ENDPOINT="blah.com"
bin/cli-template --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
unset API_KEY
unset ENDPOINTS
echo ""