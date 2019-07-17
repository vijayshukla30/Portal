# User Service
## Crete a Proto File
user/proto/account.proto

## Generate Go File
```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. srv/proto/hello.proto
```

## Create User Service and Main
user/internal and user/cmd.
Service Name is go.micro.srv.user

## Run User Service
```
cd NexthoughtsPortal
go run user/cmd/userd/main.go
 ```
# API
## Create a Service and Main
api/internal and user/cmd
Service Name is go.micro.api.v1.nexthoughts

URL will be http://localhost:8080/v1/nexthoughts/

## Run API Service
```
cd NexthoughtsPortal
go run api/cmd/apid/main.go
```

# Run Micro API with HTTP
```
micro api --handler=http
```

## Create User using
http://localhost:8080/v1/nexthoughts/users/create