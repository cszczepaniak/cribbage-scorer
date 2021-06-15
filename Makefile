.PHONY: build-go
build-go:
	cd go && go build -o ../go.exe

.PHONY: build-rust
build-rust:
	cd rust && cargo build