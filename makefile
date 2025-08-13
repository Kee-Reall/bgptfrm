BINARY_NAME=bgptfrm
OUT_DIR=./out


build: clean
	mkdir -p $(OUT_DIR)
	CGO_ENABLED=0 GOOS=linux go build -o $(OUT_DIR)/$(BINARY_NAME) ./main.go

clean:
	rm -rf $(OUT_DIR)