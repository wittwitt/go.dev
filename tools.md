# tools.md

```sh
go install xxx@latest 可以之直接安装，不需要执行环境是有go.mod

```

## modgraphviz

golang.org/x/exp/cmd/modgraphviz

```sh

go install golang.org/x/exp/cmd/modgraphviz@latest

go mod graph | modgraphviz > graph.dot
go mod graph | modgraphviz | dot -Tpng -o graph.png
```

### proto

```sh
// proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install github.com/gogo/protobuf/protoc-gen-gogo@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
helloworld.proto
```

### wrie

```sh

go install github.com/google/wire/cmd/wire@latest

cd /path/to/wire_provieder.go

wire

```