build:
	CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/memstress main.go
