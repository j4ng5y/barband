PKG_DB := pkg/db
PKG_SERVER_BUFFS := pkg/server/buffs
SQL_FILES := $(shell find . -type f -name *.sql)
PROTO_FILES := $(shell ls docs/protobufs)

.PHONY: generate
generate : go_db_files proto_bindings

# Generate the pkg/db directory
$(PKG_DB) :
	@mkdir -p $@

# Generate the pkg/server/buffs directory
$(PKG_SERVER_BUFFS) :
	@mkdir -p $@

.PHONY: go_db_files
# Generate SQL code
go_db_files : $(PKG_DB) $(SQL_FILES)
	@echo "Generating SQL bindings"
	@sqlc generate

.PHONY: proto_bindings
# Generate Protobuf bindings
proto_bindings : $(PKG_SERVER_BUFFS)
	@echo "Generating gRPC bindings"
	@protoc -I docs/protobufs \
	--go_out=$(PKG_SERVER_BUFFS) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PKG_SERVER_BUFFS) \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_FILES)