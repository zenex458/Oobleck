scripts=example_script

all: $(scripts) build/Oobleck
.PHONY: all

# Preparing build directory
build:
	mkdir -p build/scripts

build/Oobleck: main.go
	go build -o build/Oobleck main.go

$(scripts): # TODO: this is inefficient, I'm not knowledgable in make, but this recompiles everytime and it should only recompile when changes in $@.go are made
	go build -buildmode=plugin -o build/scripts/$@.so scripts/$@.go

clean:
	rm -rf build
