all: build  

build:
	go build -o api.com ./main.go

clean:
	go clean -i ./...

test:
	go test -race ./...

