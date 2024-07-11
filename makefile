TARGET := cfs

all: clean build

build:
	mkdir -p build
	cd cmd/config_switcher/ && go build -o ../../build/$(TARGET) .

run:
	./build/$(TARGET)

clean:
	touch *
	rm -rf build
