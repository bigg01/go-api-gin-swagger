get_deps:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest

update_go_deps:
	go get -u
	go mod tidy

swag_init:
	swag init
	swag fmt

hello:
	echo "Hello"

build:
	go build -o bin/trinity_server main.go
build_native:
	#GOOS=linux GOARCH=amd64 go build -o bin/trinity_server -a -installsuffix cgo main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o bin/trinity_server -a -installsuffix cgo main.go
compile:
	echo "Compiling for every OS and Platform"
	# -a -installsuffix cgo
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build  -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go


upx:
	echo "---> upx compress "
	du -hs bin/trinity_server
	upx -qq bin/trinity_server
	echo "---> upx verify "
	upx -t bin/trinity_server
	du -hs bin/trinity_server

container:
	docker build .  --file Dockerfile --tag trinity_server:latest

containerrun:
	docker run -p 8080:8080 trinity_server:latest

buildcontainer: build_native upx container


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

prom_start:
	podman-compose -f compose/docker-compose-prom.yml up
prom_stop:
	podman-compose -f compose/docker-compose-prom.yml down
prom_reload:
	curl -X POST localhost:9090/-/reload

cosmosdb_up:
	podman-compose -f compose/docker-compose-cosmosdb.yml up

cosmosdb_stop:
	podman-compose -f compose/docker-compose-cosmosdb.yml down