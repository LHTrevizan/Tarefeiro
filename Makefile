APP_NAME= tarafeiro
VERSION= 0.1.0

CMD_PATH = ./cmd/$(APP_NAME)

.PHONY : run build test clean

run :
	go run $(CMD_PATH)
build :
	CGO_ENABLED=0 go build -o $(BINARY_NAME) ./cmd/tarefeiro

test :
	go test ./...

clean :
	rm -rf $(APP_NAME)