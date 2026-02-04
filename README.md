# bunnymark

Benchmark test for go module [github.com/bernhardfritz/ecs](https://github.com/bernhardfritz/ecs).

## Setup

```
git clone https://github.com/bernhardfritz/bunnymark.git
git clone https://github.com/BrownNPC/Raylib-Go-Wasm.git
cd Raylib-Go-Wasm
git apply ../bunnymark/server.patch
cd ../bunnymark
docker run -it --rm -v $PWD:/app -v $PWD/../Raylib-Go-Wasm:/app/Raylib-Go-Wasm --workdir /app golang:1.25.6-alpine3.23 sh
go build ./Raylib-Go-Wasm/server/server.go
go get -tool github.com/air-verse/air@latest
```

## Development

```
docker run -it --rm -v $PWD:/app -v $PWD/../Raylib-Go-Wasm:/app/Raylib-Go-Wasm --workdir /app golang:1.25.6-alpine3.23 sh
go tool air
./server # run in a separate terminal
```

## References

* https://github.com/gen2brain/raylib-go
* https://github.com/BrownNPC/Raylib-Go-Wasm
* https://github.com/air-verse/air
* https://github.com/aarol/reload