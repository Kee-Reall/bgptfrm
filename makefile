BINARY_NAME=bgptfrm
OUT_DIR=./out
MAIN_PATH=


build: clean
	mkdir -p $(OUT_DIR)
	CGO_ENABLED=0 GOOS=linux go build -o $(OUT_DIR)/$(BINARY_NAME) ./src/main.go

clean:
	rm -rf $(OUT_DIR)/*

run:
	go run ./src/main.go