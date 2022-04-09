get_deps:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest

swag_init:
	swag init
	swag fmt

hello:
	echo "Hello"

build:
	go build -o bin/trinity_server main.go
compile:
	echo "Compiling for every OS and Platform"
	# -a -installsuffix cgo
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build  -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go


upx:
	echo "---> upx compress "
	upx -qq bin/main
	echo "---> upx verify "
	upx -t bin/main

container:
	docker build . --file Dockerfile --tag mycontainer:latest

containerrun:
	docker run mycontainer:latest

buildcontainer: build container


coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out

test:
	go test -v

sec:
	gosec -tests ./...
#gosec -fmt=json -out=results.json ./...

run:
	go run main.go

run_bin:
	./bin/trinity_server

clean:
	rm -rf bin/*

all: hello build