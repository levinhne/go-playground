.PHONY: generate

BUF_VERSION:=v1.30

generate: generate/proto
generate/proto:
	cd ./proto && go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) mod update
	@go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate
