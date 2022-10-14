# svcgen

This tool is designed to generate an empty microservice [blanksvc](https://github.com/ceperapl/blanksvc)

## Run and test

```bash
# from svcgen folder:
go run main.go --service.name=todo --service.path=/home/siarhei_pazdniakou --module.path=github.com/ceperapl/todo

# from todo folder:
task generate
go mod tidy
task build
./bin/todo
#> ts=2022-10-14T12:04:51.207503784Z caller=server.go:50 transport=HTTP addr=0.0.0.0:8080
#> ts=2022-10-14T12:04:51.208141289Z caller=server.go:65 transport=GRPC addr=0.0.0.0:9000

```

The structure of created microservice:

```
todo
├── api
│   └── protobuf-spec
│       ├── hello/v1/hello.proto
│       └── buf.yaml
├── bin
│   └── todo
├── cmd
│   ├── config.go
│   └── server.go
├── gen/proto/go/hello/v1
│                      ├── hello_grpc.pb.go
│                      └── hello.pb.go
├── pkg
│   ├── endpoints
│   │   ├── middleware/logging.go
│   │   └── endpoints.go
│   ├── service
│   │   ├── errors.go
│   │   └── service.go
│   ├── transport
│   │   ├── grpc/grpc.go
│   │   └── http
│   │       ├── errors.go
│   │       └── http.go
│   └── utils/healthcheck/healthcheck.go
├── buf.gen.yaml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
└── Taskfile.yml
```