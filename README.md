https://github.com/swaggo/swag

swag init

swag fmt

http://localhost:8080/swagger/index.html


https://docs.microsoft.com/en-us/azure/cosmos-db/table/how-to-use-go?tabs=bash



https://github.com/go-swagger/go-swagger/blob/master/docs/generate/client.md


https://github.com/go-swagger/go-swagger

# generate cli
```bash
swagger generate cli -f ../docs/swagger.json --cli-app-name trinityclient
go get -u -f ./...
go build -o trinityclient cmd/trinityclient/main.go
```

```bash
./trinityclient
Usage:
trinityclient [command]

Available Commands:
album       
albums      
completion  Generate completion script
example     
help        Help about any command

Flags:
--config string     config file path
--debug             output debug logs
--dry-run           do not send the request to server
-h, --help              help for trinityclient
--hostname string   hostname of the service (default "localhost:8080")
--scheme string     Choose from: [http] (default "http")

Use "trinityclient [command] --help" for more information about a command.
```

```bash
./trinityclient --hostname localhost:8080 album GetAlbAlbums
[{"artist":"John Coltrane","id":"1","price":56.99,"title":"Blue Train"},{"artist":"Gerry Mulligan","id":"2","price":17.99,"title":"Jeru"},{"artist":"Sarah Vaughan","id":"3","price":39.99,"title":"Sarah Vaughan and Clifford Brown"},{"artist":"Oliver Gugenbuehl","id":"4","price":1000.99,"title":"The Masters of Cloud"}]

```
# go-api-gin-swagger
