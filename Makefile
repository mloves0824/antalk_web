.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: antalk-web-fe 


antalk-web-fe: 
	go build -i -o bin/antalk-web-fe ./cmd/fe

clean: 
	@rm -rf bin
