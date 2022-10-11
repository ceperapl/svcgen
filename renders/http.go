package renders

import (
	. "github.com/dave/jennifer/jen"
)

func (svc ServiceData) RenderHttp() *File {
	file := NewFile("http")

	return file
}
