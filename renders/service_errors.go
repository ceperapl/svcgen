package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderServiceErrors() *File {
	file := NewFilePathName("errors", "service")

	return file
}
