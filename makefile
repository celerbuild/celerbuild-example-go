.PHONY: build package run-server test

test:
	@echo "Running unit tests..."
	go test ./tests/... -v

build: test
	@echo "Building the web server for Linux..."
	GOOS=linux GOARCH=amd64 go build -o ./build/server main.go
	@echo "Packaging the application..."
	cp -r shell/* ./build/
	chmod +x ./build/*.sh
	cd build && tar -zcvf ${LOCAL_PACKAGE_FILE} *

run-server:
	go run main.go