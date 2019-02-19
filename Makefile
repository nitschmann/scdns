GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./build/build-executable.sh

test:
	$(GOTEST) -v ./...

build-all:
	$(LOCAL_BUILD) linux
	$(LOCAL_BUILD) darwin
	$(LOCAL_BUILD) windows

build-linux:
	$(LOCAL_BUILD) linux

build-darwin:
	$(LOCAL_BUILD) darwin

build-windows:
	$(LOCAL_BUILD) windows

