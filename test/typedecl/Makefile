regenerate:
	go install github.com/ypenghui6/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. --proto_path=../../../../../:../../protobuf/:. typedecl.proto
