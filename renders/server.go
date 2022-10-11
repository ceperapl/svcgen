package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderServer() *File {
	file := NewFilePathName("server", "cmd")

	return file
}
