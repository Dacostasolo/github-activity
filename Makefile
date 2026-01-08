.PHONY: run build clean install
run:
	go run ./cmd ${ARGS}

build:
	go build -o bin/github-activity ./cmd

clean:
	rm -rf bin/

install:
	go install ./cmd

