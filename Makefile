# go_study — Makefile
#
# Common targets:
#   make            # build the binary
#   make run        # build and serve on :8080 (migrates automatically)
#   make sqlc       # regenerate sqlc store code from queries.sql
#   make clean      # remove binary and the local sqlite db
#
# Migrations: run automatically by `serve` on startup. There is no manual
# migrate command — the binary owns its schema. To force a re-migration from
# scratch, `make reset-db`.

BINARY    := go_study
ADDR      ?= :8080
DB        ?= go_study.db
ROOT      ?= .

.DEFAULT_GOAL := build

## help: list targets
.PHONY: help
help:
	@awk 'BEGIN{FS=":.*##"} /^## / {sub(/^## /,""); split($$0,a,":"); printf "  \033[36m%-18s\033[0m %s\n", a[1], a[2]}' $(MAKEFILE_LIST)

## build: compile the server binary
.PHONY: build
build:
	go build -o $(BINARY) .

## run: build and start the server
.PHONY: run
run: build
	./$(BINARY) serve --addr $(ADDR) --db $(DB) --root $(ROOT)

## serve: run without rebuilding (alias for `go run .`)
.PHONY: serve
serve:
	go run . serve --addr $(ADDR) --db $(DB) --root $(ROOT)

## sqlc: regenerate internal/db/store from queries.sql
.PHONY: sqlc
sqlc:
	@command -v sqlc >/dev/null 2>&1 || { echo "sqlc not installed: brew install sqlc OR go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"; exit 1; }
	sqlc generate

## test: run unit tests
.PHONY: test
test:
	go test ./internal/...

## vet: run go vet on the server code
.PHONY: vet
vet:
	go vet . ./internal/...

## fmt: gofmt the server code
.PHONY: fmt
fmt:
	gofmt -w -s . internal/

## tidy: clean up go.mod / go.sum
.PHONY: tidy
tidy:
	go mod tidy

## tools: install sqlc into $$GOPATH/bin
.PHONY: tools
tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## clean: remove binary and the local sqlite database
.PHONY: clean
clean:
	rm -f $(BINARY) $(DB) $(DB)-wal $(DB)-shm

## reset-db: drop and recreate the local database
.PHONY: reset-db
reset-db: clean build
	./$(BINARY) migrate up --db $(DB)
