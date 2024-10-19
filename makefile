.PHONY: .test

test:
	@echo "Running tests..."
	grc go test -v -vet=off ./...

run: 
	@echo "Build and running..."
	go run main.go