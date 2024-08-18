BINARY_NAME=gopass
INSTALL_DIR=/usr/local/bin
SOURCE_FILE=cmd/main.go

build:
	go build -o $(BINARY_NAME) $(SOURCE_FILE)

install: build
	sudo mv $(BINARY_NAME) $(INSTALL_DIR)/

uninstall:
	sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)

.PHONY: build install uninstall
