GO := go

.PHONY: build
build:
	$(GO) build -ldflags "-s -w" -trimpath -o bin/beershop cmd/server/main.go

.PHONY: clean
clean:
	rm -rf bin/