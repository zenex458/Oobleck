scripts = example_script

all: $(scripts) build/Oobleck
.PHONY: all

# Preparing build directory
build:
	mkdir -p build/scripts

build/Oobleck: 
	go build -o build/Oobleck src/main.go

$(scripts):
	go build -buildmode=plugin -o build/scripts/$@.so src/scripts/$@.go 

clean:
	rm -rf build
