# Advent of Code 2021

[Advent of Code 2021](https://adventofcode.com/2021)

```shell

go() {
    docker run -it -v `pwd`:/go/src/app -w /go/src/app -e GOCACHE='/go/src/app/.cache' -u `id -u` golang:1.17-bullseye go $@
}

# day1 -> day7
cd dayX
go build && ./today

# day8+

go run ./dayX
```