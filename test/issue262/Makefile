regenerate:
	go install github.com/ypenghui6/protobuf/protoc-min-version
	go install github.com/ypenghui6/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/ypenghui6/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/ypenghui6/protobuf/types:. timefail.proto
