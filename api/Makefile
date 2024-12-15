APIDIR := code
PROTODIR := proto
PROTOC := protoc

PROTO := $(shell find proto -type f -name '*.proto')
TARGET := $(patsubst $(PROTODIR)/%, $(APIDIR)/%, $(PROTO:.proto=_api))

all: $(TARGET)
	@echo $(TARGET)

$(APIDIR)/%_api: $(PROTODIR)/%.proto
	$(PROTOC) --go_out=. --go-grpc_out=. $<

setup:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: all setup
