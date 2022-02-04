GOX := $(shell which go)
BIN := puckd
PREFIX := /usr

puckd:
	$(GOX) build \
		-o $(BIN) \
		-x
clean:
	@rm -rf vendor
	@rm -f $(BIN)