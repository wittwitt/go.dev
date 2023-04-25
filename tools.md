# tools.md


```sh
go install xxx@last 可以之直接安装，不需要执行环境是有go.mod

```


## modgraphviz

golang.org/x/exp/cmd/modgraphviz

```sh

go install  golang.org/x/exp/cmd/modgraphviz@last

go mod graph | modgraphviz > graph.dot
go mod graph | modgraphviz | dot -Tpng -o graph.png
```