regenerate:
	go install github.com/ypenghui6/protobuf/protoc-gen-gogo
	go install github.com/ypenghui6/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mgoogle/protobuf/any.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/api.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/empty.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/field_mask.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/source_context.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/type.proto=github.com/ypenghui6/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/ypenghui6/protobuf/types,\
	:. --proto_path=../../../../../:../../protobuf/:. issue438.proto
