# Advent of Code 2021

[Advent of Code 2021](https://adventofcode.com/2021)

```shell
go() {
    docker run -it -v `pwd`:/go/src/app -w /go/src/app -e GOCACHE='/go/src/app/.cache' -u `id -u` golang:1.17-bullseye go $@
}

go build && ./today

```