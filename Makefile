PROTOC=protoc
PROTO_DIR=proto
AUTH_OUT_DIR=authservice/internal
AUTH_PROTO_FILE=$(PROTO_DIR)/authservice/auth.proto
AUTH_GEN_GO=$(AUTH_OUT_DIR)/auth.pb.go

auth-proto: $(AUTH_GEN_GO)

$(AUTH_GEN_GO): $(AUTH_PROTO_FILE)
	@echo "Generating authservice proto"
	@mkdir -p $(AUTH_OUT_DIR)
	$(PROTOC) --proto_path=$(PROTO_DIR) \
		--go_out=$(AUTH_OUT_DIR) \
		--go-grpc_out=$(AUTH_OUT_DIR) \
		$(AUTH_PROTO_FILE)
