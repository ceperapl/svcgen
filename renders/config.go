package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderConfig() *File {
	file := NewFilePathName("config", "cmd")

	return file
}
