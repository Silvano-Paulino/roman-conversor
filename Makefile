
tests: 
	go test -v ./internal/converter/*

build-linux: 
	GOOS=linux GOARCH=amd64 go build -o cli ./internal/main.go

build-windows: 
	GOOS=windows GOARCH=amd64 go build -o cli.exe ./internal/main.go