CONF_PROTO_FILES=$(shell find ./protobuf/conf/ -name *.proto)

.PHONY: conf
# generate conf proto
conf:
		protoc --proto_path=. \
			   --go_out=paths=source_relative:. \
			   $(CONF_PROTO_FILES)

