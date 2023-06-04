# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
BINARY_NAME = anyjson2csv

# Build directory structure
BUILD_DIR = build
LINUX_DIR = $(BUILD_DIR)/linux
WINDOWS_DIR = $(BUILD_DIR)/windows
MACOS_DIR = $(BUILD_DIR)/macos

# Targets
all: clean test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Cross compilation
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_DIR)/$(BINARY_NAME) -v

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_DIR)/$(BINARY_NAME).exe -v

build-macos:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(MACOS_DIR)/$(BINARY_NAME) -v

# Default target
.DEFAULT_GOAL := all
