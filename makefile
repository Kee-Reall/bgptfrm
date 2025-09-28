BINARY_NAME=bgptfrm
OUT_DIR=./out
MAIN_PATH=./src/main.go


build: clean
	mkdir -p $(OUT_DIR)
	CGO_ENABLED=0 GOOS=linux go build -o $(OUT_DIR)/$(BINARY_NAME) $(MAIN_PATH)

start: build
	$(OUT_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(OUT_DIR)/*

run:
	go run $(MAIN_PATH)