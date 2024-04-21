# Fundamentals of Security

## Requirements

- VsCode
- Docker

## Instructions

Clone the repo

```sh
git clone https://github.com/k1910177/funsec-report1
```

Open project in vscode

```sh
code funcsec-report1
```

Open workspace in devcontainer and wait for it to boot up

Run benchmark in integrated terminal

```sh
go test -bench . -benchmem -benchtime=1000000x -cpu=1
```

Create memory profile and view in the browser

```sh
go test -bench=BenchmarkAES128SBox . -memprofile mem.out -o pprof.bin
go tool pprof -http=":8888" mem.out
```

