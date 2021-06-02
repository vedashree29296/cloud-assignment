.PHONY: build clean 

build: 
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler main.go

clean:
	rm -rf ./bin ./vendor