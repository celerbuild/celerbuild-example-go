.PHONY: build package run-server test clean

test:
	@echo "Running unit tests..."
	go test ./tests/... -v

build: test
	@echo "Building the web server for Linux..."
	GOOS=linux GOARCH=amd64 go build -o ./build/server main.go
	@echo "Packaging the application..."
	cp -r shell/* ./build/
	chmod +x ./build/*.sh

clean:
	@echo "Cleaning up..."
	rm -f ./build/server
	rm -f ./build/server.pid
	rm -rf ./build/*.sh