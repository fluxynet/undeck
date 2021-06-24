.PHONY: make

make: clean build/undeck

build/undeck:
	@cd cmd && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -extldflags '-static'" -o ../build/undeck

clean:
	@rm -f build/undeck
