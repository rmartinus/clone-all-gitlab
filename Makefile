help: Makefile
	@echo " Choose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## run: Run clone all
run:
	go run cmd/clone-all.go
