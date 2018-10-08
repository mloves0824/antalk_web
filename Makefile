.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: antalk-web-fe test 


antalk-web-fe: 
	go build -i -o bin/antalk-web-fe ./cmd/fe
test:
	go build -i -o bin/test ./cmd/test/pegasus

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-web-image .
