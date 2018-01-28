package typedeclimport

import subpkg "github.com/Cyinx/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
