regenerate:
	go install github.com/ypenghui6/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/ypenghui6/protobuf/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/ypenghui6/protobuf/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
