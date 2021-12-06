regenerate:
	go install github.com/ypenghui6/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/ypenghui6/protobuf/protobuf/ \
	--gogo_out=. events.proto 
