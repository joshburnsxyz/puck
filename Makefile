GOX := $(shell which go)
BIN := puck
PREFIX := /usr

puck: vendor
	$(GOX) build \
		-o $(BIN) \
		-x
clean:
	@rm -rf vendor
	@rm -f $(BIN)

vendor:
	go mod vendor

install:
	cp -f $(BIN) $(PREFIX)/bin/$(BIN)

uninstall:
	rm -f $(PREFIX)/bin/$(BIN)


.PHONY: clean install