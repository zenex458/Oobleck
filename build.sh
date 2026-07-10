rm -rf ./build
mkdir -p build/tests

go build -buildmode=plugin -o build/tests/test.so src/testr.go
go build -o build/Oobleck src/main.go
