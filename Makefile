GO=go

all: build test

build: amd64 arm64

amd64: powerdns-remote-backend-linux-amd64 powerdns-remote-backend-darwin-amd64 powerdns-remote-backend-windows-amd64.exe

arm64: powerdns-remote-backend-linux-arm64 powerdns-remote-backend-darwin-arm64 powerdns-remote-backend-windows-arm64.exe

powerdns-remote-backend-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o powerdns-remote-backend-linux-amd64

powerdns-remote-backend-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o powerdns-remote-backend-darwin-amd64

powerdns-remote-backend-windows-amd64.exe:
	env GOOS=windows GOARCH=amd64 go build -o powerdns-remote-backend-windows-amd64.exe

powerdns-remote-backend-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o powerdns-remote-backend-linux-arm64

powerdns-remote-backend-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o powerdns-remote-backend-darwin-arm64

powerdns-remote-backend-windows-arm64.exe:
	env GOOS=windows GOARCH=arm64 go build -o powerdns-remote-backend-windows-arm64.exe

test:
	govet-report.json test-report.json coverage.out

lint:
	go help lint

govet-report.json:
	go vet -json ./... > govet-report.json

test-report.json:
	go test -json ./... > test-report.json

coverage.out:
	go test -coverprofile=coverage.out ./...

clean:
	rm -Rf powerdns-remote-backend-*
