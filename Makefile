#!/usr/bin/make -f

compress: compile
	zip -r 'bin-$(shell git describe).zip' ./bin

compile: prepare
	GOOS=darwin go test -c -o ./bin/macos/smartyping
	GOOS=linux GOARCH=386 go test -c -o ./bin/linux-386/smartyping
	GOOS=linux GOARCH=amd64 go test -c -o ./bin/linux-amd64/smartyping
	GOOS=windows GOARCH=386 go test -c -o ./bin/windows-386/smartyping
	GOOS=windows GOARCH=amd64 go test -c -o ./bin/windows-amd64/smartyping

prepare: clean
	mkdir -p ./bin/macos
	mkdir -p ./bin/linux-386
	mkdir -p ./bin/linux-amd64
	mkdir -p ./bin/windows-386
	mkdir -p ./bin/windows-amd64

clean:
	rm -rf ./bin*
