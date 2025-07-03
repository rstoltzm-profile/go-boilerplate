Write-Host "`n## 1. test no arguments ##`n"
& bin/cli-template.exe

Write-Host "`n## 2. test no env api_key ##`n"
& bin/cli-template.exe --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"

Write-Host "`n## 3. test no env endpoint ##`n"
$env:API_KEY = "blah"
& bin/cli-template.exe --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
Remove-Item Env:API_KEY

Write-Host "`n## 4. test cli ##`n"
$env:API_KEY = "blah"
$env:ENDPOINT = "blah.com"
& bin/cli-template.exe --variable "test" --inputFile "file.json" --outFormat "json" --outFile "outFile.json"
Remove-Item Env:API_KEY
Remove-Item Env:ENDPOINT
Write-Host ""