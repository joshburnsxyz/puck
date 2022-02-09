GOX := $(shell which go)
BIN := puck
PREFIX := /usr

puck: modinit
	$(GOX) build ./cmd \
		-o $(BIN) \
		-x
clean:
	@rm -rf vendor
	@rm -f $(BIN)

vendor:
	go mod vendor
	
modinit:
	@rm -rf vendor
	go mod tidy
	go mod vendor

install:
	cp -f $(BIN) $(PREFIX)/bin/$(BIN)

uninstall:
	rm -f $(PREFIX)/bin/$(BIN)


.PHONY: clean install modinit
