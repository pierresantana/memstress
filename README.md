# Simple memory stress test

This utility is a simple memory stress test written in go. It fills memory with a number of bytes for a period of time.

## Build

### Using the Makefile:

```bash
make build
```

## Install:

```bash
go install github.com/pierresantana/memstress@latest
```

## Usage

```bash
$ memstress -h
Usage of memstress:
  -size int
    	size in MB (default 1024)
  -timeout duration
    	timeout to release the memory (default 1m0s)

$ memstress -size 2048 -timeout 2m
2022/10/27 09:50:41 Filling the memory with 2048 MB
2022/10/27 09:50:54 Memory filled. Waiting 2m0s
2022/10/27 09:52:54 Releasing memory
```
