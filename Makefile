scripts=$(shell find scripts/ -type f -exec basename {} .go \;)
buildplugins=$(foreach wrd,$(scripts),go build -buildmode=plugin -o build/scripts/$(wrd).so scripts/$(wrd).go;)
buildpluginsrelease=$(foreach wrd,$(scripts),go build -ldflags '-w -s' -buildmode=plugin -o build/scripts/$(wrd).so scripts/$(wrd).go;)
scriptsobjects=$(foreach wrd,$(scripts),$(wrd).so)


.PHONY: run release clean

build/scripts/$(scriptsobjects): build/Oobleck
	$(buildplugins)

build/Oobleck: main.go
	go build -o build/Oobleck main.go

run: build/scripts/$(scripts).so
	./build/Oobleck

release:
	go build -ldflags '-w -s' -o build/Oobleck main.go
	GOOS=windows go build -ldflags '-w -s' -o build/Oobleck.exe main.go
	$(buildpluginsrelease)
#todo: buildplugins for windows
clean:
	rm -rf build
