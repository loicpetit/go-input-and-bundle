Experiment running go module with input stream

The file to read is in input folder

The module must manage the cases : the file is missing, wrong or correct

Run the module with a file : Get-Content ..\files\input.txt | go run .\main.go

Bundle
go install golang.org/x/tools/cmd/bundle@latest
& ($env:USERPROFILE + "/go/bin/bundle") -o ../dist/main.go example/input-and-bundle
Bundle only one package...
& ($env:USERPROFILE + "/go/bin/bundle") -o ../dist/tobundle.go example/input-and-bundle/tobundle  
