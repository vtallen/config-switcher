TARGET := cfs
INSTALL_PATH := /usr/bin/
CONFIG_FOLDER := ~/.config/$(TARGET)/

all: clean build

build:
	mkdir -p build
	cd cmd/config_switcher/ && go build -o ../../build/$(TARGET) .

run:
	./build/$(TARGET)

install: clean build
	sudo mkdir -p $(CONFIG_FOLDER)
	sudo mv config.yaml $(CONFIG_FOLDER)
	sudo mv build/$(TARGET) $(INSTALL_PATH)
		
clean:
	touch *
	rm -rf build
