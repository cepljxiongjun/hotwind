BINARY_NAME=hotwind
run:
	go build -o bin/$(BINARY_NAME) -v
	./bin/$(BINARY_NAME)