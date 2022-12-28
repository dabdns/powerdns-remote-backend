GO=go

all: build report.json coverage.out

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

lint:
	go help lint

vet:
	go vet

report.json:
	go test -json ./... > report.json

coverage.out:
	go test -coverprofile=coverage.out ./...

clean:
	rm -Rf powerdns-remote-backend-*
