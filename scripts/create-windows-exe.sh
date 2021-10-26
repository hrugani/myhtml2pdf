# build the .exe binary
cd ../cmd/webapi
GOOS=windows go build -o ../../dist/windows-dist/mypdfservices.exe main.go
cd ../../scripts