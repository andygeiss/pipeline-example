TS=$(shell date -u '+%Y/%m/%d %H:%M:%S')

all: generate_api run

generate_api:
	@echo $(TS) Generating API ...
	@protoc --go_out=. internal/api/*.proto
	@echo $(TS) Done.

run:
	@echo $(TS) Running ...
	@go run cmd/pipeline/main.go
	@echo $(TS) Done.

setup:
	@echo $(TS) Installing dependencies ...
	@sudo apt-get install -y protobuf-compiler
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
	@echo $(TS) Done.
