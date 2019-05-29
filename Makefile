help: Makefile
	@echo " Choose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: build binary version in bin/myclone
build:
	go build -o bin/clone-all cmd/clone-all.go

## run: Run clone all
run:
	go run cmd/clone-all.go
