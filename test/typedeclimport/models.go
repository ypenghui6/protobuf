package typedeclimport

import subpkg "github.com/ypenghui6/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
