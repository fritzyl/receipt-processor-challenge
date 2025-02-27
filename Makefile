BINARY_NAME = receipt-processor

# Define the build target
build:
	go build -o $(BINARY_NAME) .

# Define the run target
run: build
	./$(BINARY_NAME)

# Define the clean target
clean:
	go clean
	rm -f $(BINARY_NAME)

# Define the test target
test:
	go test ./api/rules

# Default target
.DEFAULT_GOAL := build